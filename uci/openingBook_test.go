package uci

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestFirstMove(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	m := make(map[string]int)

	const runs = 10_000

	for i := 0; i < runs; i++ {
		uci := getFirstMove()
		m[uci] += 1
	}

	type a struct {
		uci  string
		freq int
	}

	var list []a

	for k, v := range m {
		list = append(list, a{uci: k, freq: v})
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].freq > list[j].freq
	})

	for _, item := range list {
		fmt.Printf("%s: %4d %4.1f%%\n", item.uci, item.freq, float64(item.freq)/runs*100)
	}
}
