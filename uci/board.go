package uci

import (
	"strings"
	"unicode"
)

type Board struct {
	Pos             []rune
	ActiveColor     string
	Castling        string
	EnPassantSquare string
	HalfmoveClock   string
	FullMove        string
}

func FENtoBoard(fen string) Board {
	parts := strings.Split(fen, " ")
	ranks := strings.Split(parts[0], "/")
	b := Board{
		ActiveColor:     parts[1],
		Castling:        parts[2],
		EnPassantSquare: parts[3],
		HalfmoveClock:   parts[4],
		FullMove:        parts[5],
		Pos:             make([]rune, 64),
	}

	for i := 7; i >= 0; i-- {
		rank := ranks[i]
		offset := i * 8
		for _, c := range rank {
			if unicode.IsDigit(c) {
				n := int(c) - 48
				for j := 0; j < n; j++ {
					b.Pos[offset] = ' '
					offset++
				}
			} else {
				b.Pos[offset] = c
				offset++
			}
		}
	}

	return b
}
