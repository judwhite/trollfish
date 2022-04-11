package main

import (
	"context"
	"runtime"
	"trollfish/uci"
)

// UCI stuff:

// setoption name <name> value <value>
// - if 'button' no value needed (toggle?)
// - option names are not case sensitive
// examples:
// "setoption name Selectivity value 3\n"
// "setoption name Clear Hash\n"

// expanded algebraic notation:
// - Examples: e2e4, e7e5, e1g1 (white short castling), e7e8q (for promotion)
// - nullmove = 0000

// -> debug [on | off]
// <- info string <xxxxx>

// -> cmd1
// -> cmd2
// -> cmd3
// -> isready
// <- readyok (when commands complete, ex setoption for tablebases)

// ucinewgame - sent btwn games
// position [fen | startpos ] moves ...

// go depth ##
// go movetime ### (milliseconds?)
// go infinite (run until stop)
// stop

func main() {
	p := uci.New("yancy 0.01", "the yancy developers", uci.Option{Name: "threads", Type: uci.OptionTypeSpin, Default: "1", Min: 1, Max: runtime.NumCPU()})
	ctx, _ := p.Start(context.Background())
	<-ctx.Done()
}
