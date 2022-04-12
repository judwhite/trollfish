package stockfish

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

type StockFish struct {
	Ctx    context.Context
	Output <-chan string

	cancel  context.CancelFunc
	writer  io.WriteCloser
	logInfo func(string)
}

func Start(ctx context.Context, binary string, logInfo func(string)) (*StockFish, error) {
	_, err := os.Stat(binary)
	if err != nil && os.IsNotExist(err) {
		return nil, fmt.Errorf("'%s' not found", binary)
	}

	dir := filepath.Base(binary)

	output := make(chan string, 512)

	var sf StockFish
	sf.Ctx, sf.cancel = context.WithCancel(ctx)
	sf.Output = output
	sf.logInfo = logInfo

	cmd := exec.CommandContext(sf.Ctx, binary)
	cmd.Dir = dir

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	sf.writer = stdin

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		log.Fatalf("cmd.Start: %v\n", err)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	// stderr loop
	go func() {
		defer wg.Done()
		r := bufio.NewScanner(stderr)
		for r.Scan() {
			select {
			case <-sf.Ctx.Done():
				return
			default:
				line := r.Text()
				logInfo(fmt.Sprintf("SF STDERR: %s", line))
			}
		}
		if err := r.Err(); err != nil {
			logInfo(fmt.Sprintf("SF ERR: stderr: %v", err))
		}
	}()

	// stdout loop
	go func() {
		defer wg.Done()
		r := bufio.NewScanner(stdout)
		for r.Scan() {
			select {
			case output <- r.Text():
			case <-sf.Ctx.Done():
				return
			}
		}
		if err := r.Err(); err != nil {
			logInfo(fmt.Sprintf("SF ERR: stdout: %v", err))
		}
	}()

	go func() {
		if err := cmd.Wait(); err != nil {
			logInfo(fmt.Sprintf("SF ERR: %v\n", err))
		}
	}()

	return &sf, nil
}

func (sf *StockFish) Write(s string) {
	sf.logInfo(fmt.Sprintf("SF: -> %s", s))

	b := []byte(s)
	b = append(b, '\n')

	_, _ = sf.writer.Write(b)
}

func (sf *StockFish) Quit() {
	sf.cancel()
}
