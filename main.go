package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"strings"
	"time"
	"trollfish/uci"
)

func createDownloadScripts() {
	urls := []string{"https://tablebase.lichess.ovh/tables/standard/6-dtz/", "https://tablebase.lichess.ovh/tables/standard/6-wdl/"}
	for i, tbURL := range urls {
		var sb strings.Builder
		resp, err := http.Get(tbURL)
		if err != nil {
			log.Fatal(err)
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()
		if resp.StatusCode != 200 {
			log.Fatalf("%s\nhttp status code %d\n", b, resp.StatusCode)
		}

		text := string(b)

		lines := strings.Split(text, "\n")
		for _, line := range lines {
			const find = "<a href=\""
			if !strings.HasPrefix(line, find) {
				continue
			}
			line = strings.TrimPrefix(line, find)
			idx := strings.Index(line, `"`)
			line = line[:idx]
			sb.WriteString(fmt.Sprintf("wget %s\n", tbURL+line))
		}

		var fileName string
		if i == 0 {
			fileName = "dtz6.sh"
		} else if i == 1 {
			fileName = "wdl6.sh"
		} else {
			log.Fatalf("index %d out of range", i)
		}

		if err := ioutil.WriteFile(fileName, []byte(sb.String()), 0644); err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	createDownloadScripts()
	return

	rand.Seed(time.Now().UnixNano())

	p := uci.New("trollfish 15", "the trollfish developers",
		uci.Option{Name: "Threads", Type: uci.OptionTypeSpin, Default: "1", Min: 1, Max: runtime.NumCPU()},
		uci.Option{Name: "MultiPV", Type: uci.OptionTypeString, Default: "8"},
		uci.Option{Name: "PlayBad", Type: uci.OptionTypeString, Default: "false"},
		uci.Option{Name: "StartAgro", Type: uci.OptionTypeString, Default: "false"},
		uci.Option{Name: "SyzygyPath", Type: uci.OptionTypeString, Default: ""},
	)
	ctx, _ := p.Start(context.Background())
	<-ctx.Done()
}
