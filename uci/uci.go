package uci

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"trollfish/lichess"
)

const startPosFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

type UCI struct {
	name    string
	author  string
	options []Option

	fen string

	started int64

	ctx    context.Context
	cancel context.CancelFunc

	mtxStdout sync.Mutex
	log       io.WriteCloser
}

func New(name, author string, options ...Option) *UCI {
	return &UCI{
		name:    name,
		author:  author,
		options: options,
	}
}

func (u *UCI) Start(ctx context.Context) (context.Context, context.CancelFunc) {
	if !atomic.CompareAndSwapInt64(&u.started, 0, 1) {
		return u.ctx, u.cancel
	}

	fp, err := os.OpenFile("yancy.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	u.log = fp

	u.ctx, u.cancel = context.WithCancel(ctx)

	c := make(chan string, 512)

	go func() {
		defer close(c)
		r := bufio.NewScanner(os.Stdin)

		for r.Scan() {
			select {
			case c <- r.Text():
			case <-u.ctx.Done():
				_ = u.log.Close()
				return
			}
		}

		if err := r.Err(); err != nil {
			msg := fmt.Sprintf("info ERR: %v", err)
			u.WriteLine(msg)
		}
	}()

	go func() {
		for line := range c {
			u.parseLine(line)
		}
	}()

	return u.ctx, u.cancel
}

func (u *UCI) parseLine(line string) {
	_, _ = u.log.Write([]byte(fmt.Sprintf("%s -> %s\n", ts(), line)))

	parts := strings.Split(strings.TrimSpace(line), " ")
	if len(parts) == 0 {
		return
	}

	switch parts[0] {
	case "uci":
		u.SetUCI()
	case "quit":
		u.Quit()
	case "isready":
		// TODO
		u.WriteLine("readyok")
	case "ucinewgame":
	// ignore
	case "setoption":
		u.setOptionRaw(parts[1:]...)
	case "position":
		u.SetPosition(parts[1:]...)
	case "stop":
		// TODO: stop any searching threads
	case "go":
		pos, err := lichess.Lookup(u.fen, "")
		if err != nil {
			u.WriteLine(fmt.Sprintf("info ERR %v", err))
			break
		}

		parentTotal := pos.White + pos.Black + pos.Draws

		start := fmt.Sprintf("info %d moves found w:%d b:%d d:%d t:%d", len(pos.Moves), pos.White, pos.Black, pos.Draws, parentTotal)
		lines := make([]string, 0, len(pos.Moves)+1)
		lines = append(lines, start)
		for _, move := range pos.Moves {
			totalGames := move.White + move.Black + move.Draws

			popularity := float64(totalGames) / float64(parentTotal)
			whiteWin := float64(move.White) / float64(totalGames)
			blackWin := float64(move.Black) / float64(totalGames)
			draw := float64(move.Draws) / float64(totalGames)

			s := fmt.Sprintf("info %5s %s %0.1f w:%d (%0.0f) b:%d (%0.0f) d:%d (%0.0f) rating:%d total:%d",
				move.SAN, move.UCI, popularity*100,
				move.White, whiteWin*100,
				move.Black, blackWin*100,
				move.Draws, draw*100,
				move.AverageRating, totalGames)
			lines = append(lines, s)
		}
		u.WriteLines(lines...)
		// TODO: handle 'infinite' and 'movetime <ms>'
	case "":
	// no-op
	default:
		msg := fmt.Sprintf("info unknown command '%s'", parts[0])
		u.WriteLine(msg)
	}
}

func (u *UCI) Quit() {
	u.cancel()
}

func (u *UCI) SetUCI() {
	var opts []string
	for _, o := range u.options {
		switch o.Type {
		case OptionTypeCheck:
		case OptionTypeSpin:
			opts = append(opts, fmt.Sprintf("option name %s type spin default %s min %d max %d", o.Name, o.DefaultValue(), o.Min, o.Max))
		case OptionTypeCombo:
		case OptionTypeButton:
		case OptionTypeString:
			opts = append(opts, fmt.Sprintf("option name %s type string default %s", o.Name, o.DefaultValue()))
		}
	}

	lines := make([]string, 0, len(opts)+4)

	lines = append(lines, fmt.Sprintf("id name %s", u.name))
	lines = append(lines, fmt.Sprintf("id author %s", u.author))
	lines = append(lines, "")
	lines = append(lines, opts...)
	lines = append(lines, "uciok")

	u.WriteLines(lines...)
}

func (u *UCI) SetOption(name, value string) {
	switch name {
	case "threads":
		n, err := strconv.Atoi(value)
		if err != nil || n < 1 {
			u.WriteLine(fmt.Sprintf("info option thread value %s invalid", value))
			return
		}
		runtime.GOMAXPROCS(n)
	default:
		u.WriteLine(fmt.Sprintf("info option %s not found", name))
	}
}

func (u *UCI) setOptionRaw(v ...string) {
	if len(v) == 0 {
		return
	}

	if v[0] != "name" {
		return
	}

	i := 1

	var name string
	for ; i < len(v); i++ {
		if v[i] == "value" {
			break
		}

		if name != "" {
			name += " "
		}
		name += v[i]
	}

	if i == len(v) || v[i] != "value" {
		// TODO: only valid for buttons
		return
	}

	var value string
	for ; i < len(v); i++ {
		if value != "" {
			value += " "
		}
		value += v[i]
	}
}

func (u *UCI) SetPosition(v ...string) {
	if len(v) == 0 {
		return
	}

	cmd := v[0]

	if cmd == "fen" {
		u.fen = strings.Join(v[1:], " ")
		u.WriteLine(fmt.Sprintf("info fen set to '%s'", u.fen))
		return
	}

	if cmd != "startpos" {
		// unknown
		u.WriteLine(fmt.Sprintf("info ERR: position '%s' command unknown", cmd))
		return
	}

	u.fen = startPosFEN
	u.WriteLine(fmt.Sprintf("info fen set to '%s'", u.fen))

	if len(v) == 1 {
		return
	}

	cmd = v[0]

	if cmd != "moves" {
		u.WriteLine(fmt.Sprintf("info ERR: position startpos '%s' command unknown", cmd))
		return
	}

	// TODO: handle moves
}

func (u *UCI) WriteLine(s string) {
	u.mtxStdout.Lock()
	defer u.mtxStdout.Unlock()
	_, _ = u.log.Write([]byte(fmt.Sprintf("%s <- %s\n", ts(), s)))
	_, _ = fmt.Fprintln(os.Stdout, s)
}

func (u *UCI) WriteLines(v ...string) {
	var w strings.Builder
	var w2 strings.Builder
	for _, s := range v {
		w.WriteString(s)
		w.WriteRune('\n')

		w2.WriteString(ts() + " <- ")
		w2.WriteString(s)
		w2.WriteRune('\n')
	}
	s := w.String()
	s2 := w2.String()

	u.mtxStdout.Lock()
	defer u.mtxStdout.Unlock()
	_, _ = u.log.Write([]byte(fmt.Sprintf("%s", s2)))
	_, _ = fmt.Fprint(os.Stdout, s)
}

func ts() string {
	return fmt.Sprintf("[%s]", time.Now().Format("2006-01-02 15:04:05"))
}
