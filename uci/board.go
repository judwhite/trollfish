package uci

import (
	"fmt"
	"math"
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

func (b *Board) FEN() string {
	var fen strings.Builder
	for i := 0; i < 8; i++ {
		if fen.Len() != 0 {
			fen.WriteRune('/')
		}

		offset := i * 8
		blanks := 0

		for j := 0; j < 8; j++ {
			if b.Pos[offset+j] == ' ' {
				blanks++
				continue
			}

			if blanks != 0 {
				fen.WriteString(fmt.Sprintf("%d", blanks))
				blanks = 0
			}

			fen.WriteRune(b.Pos[offset+j])
		}

		if blanks != 0 {
			fen.WriteString(fmt.Sprintf("%d", blanks))
			blanks = 0
		}
	}

	fen.WriteString(fmt.Sprintf(" %s %s %s %s %s", b.ActiveColor, b.Castling, b.EnPassantSquare, b.HalfmoveClock, b.FullMove))

	return fen.String()
}

func (b *Board) Moves(moves ...string) {
	halfMoveClock := atoi(b.HalfmoveClock)
	fullMove := atoi(b.FullMove)

	var activeColor int
	if b.ActiveColor == "b" {
		activeColor = 1
	}

	var wk, wq, bk, bq bool
	for _, c := range b.Castling {
		switch c {
		case 'K':
			wk = true
		case 'Q':
			wq = true
		case 'k':
			bk = true
		case 'q':
			bq = true
		}
	}

	for _, move := range moves {
		if activeColor == 1 {
			activeColor = 0
			fullMove++
		} else {
			activeColor = 1
		}

		fromUCI := move[:2]
		toUCI := move[2:]

		// castling privileges
		if fromUCI == "a1" || toUCI == "a1" {
			wq = false
		} else if fromUCI == "h1" || toUCI == "h1" {
			wk = false
		} else if fromUCI == "a8" || toUCI == "a8" {
			bq = false
		} else if fromUCI == "h8" || toUCI == "h8" {
			bk = false
		} else if fromUCI == "e1" {
			wk, wq = false, false
		} else if fromUCI == "e8" {
			bk, bq = false, false
		}

		from, to := uciToIndex(fromUCI), uciToIndex(toUCI)
		b.Pos[to] = b.Pos[from]
		b.Pos[from] = ' '

		piece := b.Pos[to]

		if toUCI == b.EnPassantSquare {
			var captureOn int
			if activeColor == 0 {
				captureOn = to - 8 // next move is white's, so the target is in black's position
			} else {
				captureOn = to + 8
			}
			b.Pos[captureOn] = ' '
		}

		// pawn move, reset halfmove clock; en passant square
		b.EnPassantSquare = "-"
		if piece == 'P' || piece == 'p' {
			halfMoveClock = 0
			if int(math.Abs(float64(to-from))) == 16 {
				var file rune
				if activeColor == 0 {
					file = '6' // next move is white's, so the target is in black's position
				} else {
					file = '3'
				}
				b.EnPassantSquare = fmt.Sprintf("%c%c", 'a'+to%8, file)
			}
		} else {
			halfMoveClock++
		}

		// white king castle
		if piece == 'K' && fromUCI == "e1" {
			if toUCI == "g1" {
				// king side
				b.Pos[to+1] = ' '
				b.Pos[to-1] = 'R'
			} else if toUCI == "c1" {
				// queen side
				b.Pos[to-2] = ' '
				b.Pos[to+1] = 'R'
			}
		}

		// black king castle
		if piece == 'k' && fromUCI == "e8" {
			if toUCI == "g8" {
				// king side
				b.Pos[to+1] = ' '
				b.Pos[to] = 'R'
			} else if toUCI == "c8" {
				// queen side
				b.Pos[to-2] = ' '
				b.Pos[to] = 'R'
			}
		}
	}

	if activeColor == 0 {
		b.ActiveColor = "w"
	} else {
		b.ActiveColor = "b"
	}

	// castling
	var cstl strings.Builder
	if wk {
		cstl.WriteRune('K')
	}
	if wq {
		cstl.WriteRune('Q')
	}
	if bk {
		cstl.WriteRune('k')
	}
	if bq {
		cstl.WriteRune('q')
	}
	b.Castling = cstl.String()

	// NOTE: en passant target square handling per move

	b.HalfmoveClock = fmt.Sprintf("%d", halfMoveClock)
	b.FullMove = fmt.Sprintf("%d", fullMove)
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

func uciToIndex(uci string) int {
	file := int(uci[0]) - 'a'
	rank := int(uci[1]) - '0' - 1
	idx := (7-rank)*8 + file
	return idx
}
