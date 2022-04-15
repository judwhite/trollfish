package main

import (
	"context"
	"runtime"
	"trollfish/uci"
)

func main() {
	p := uci.New("trollfish 14.1", "the trollfish developers",
		uci.Option{Name: "Threads", Type: uci.OptionTypeSpin, Default: "1", Min: 1, Max: runtime.NumCPU()},
		uci.Option{Name: "MultiPV", Type: uci.OptionTypeString, Default: "8"},
		uci.Option{Name: "PlayBad", Type: uci.OptionTypeString, Default: "false"},
	)
	ctx, _ := p.Start(context.Background())
	<-ctx.Done()
}
