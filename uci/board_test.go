package uci

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestFENtoBoard(t *testing.T) {
	// arrange
	cases := []struct {
		fen                 string
		wantActiveColor     string
		wantCastling        string
		wantEnPassantSquare string
		wantHalfMoveClock   string
		wantFullMove        string
		wantPos             []rune
	}{
		{
			fen:                 "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
			wantActiveColor:     "w",
			wantCastling:        "KQkq",
			wantEnPassantSquare: "-",
			wantHalfMoveClock:   "0",
			wantFullMove:        "1",
			wantPos: []rune{
				'r', 'n', 'b', 'q', 'k', 'b', 'n', 'r',
				'p', 'p', 'p', 'p', 'p', 'p', 'p', 'p',
				' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
				' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
				' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
				' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
				'P', 'P', 'P', 'P', 'P', 'P', 'P', 'P',
				'R', 'N', 'B', 'Q', 'K', 'B', 'N', 'R',
			},
		},
		{
			fen:                 "r1b1kbnr/pppp1ppp/2n5/4P3/1q6/5N2/PPPBPPPP/RN1QKB1R b KQkq - 6 5",
			wantActiveColor:     "b",
			wantCastling:        "KQkq",
			wantEnPassantSquare: "-",
			wantHalfMoveClock:   "6",
			wantFullMove:        "5",
			wantPos: []rune{
				'r', ' ', 'b', ' ', 'k', 'b', 'n', 'r',
				'p', 'p', 'p', 'p', ' ', 'p', 'p', 'p',
				' ', ' ', 'n', ' ', ' ', ' ', ' ', ' ',
				' ', ' ', ' ', ' ', 'P', ' ', ' ', ' ',
				' ', 'q', ' ', ' ', ' ', ' ', ' ', ' ',
				' ', ' ', ' ', ' ', ' ', 'N', ' ', ' ',
				'P', 'P', 'P', 'B', 'P', 'P', 'P', 'P',
				'R', 'N', ' ', 'Q', 'K', 'B', ' ', 'R',
			},
		},
	}

	sq := func(c rune) string {
		if c == 0 {
			return " "
		}
		return string(c)
	}

	for _, c := range cases {
		t.Run(c.fen, func(t *testing.T) {
			// act
			board := FENtoBoard(c.fen)

			// assert
			if !reflect.DeepEqual(c.wantPos, board.Pos) {
				var (
					loc       string
					want, got strings.Builder
				)

				want.WriteString("   abcdefgh\n   --------\n")
				got.WriteString("   abcdefgh\n   --------\n")
				for i := 0; i < 8; i++ {
					offset := i * 8
					rank := '8' - i
					want.WriteString(fmt.Sprintf("%c: ", rank))
					got.WriteString(fmt.Sprintf("%c: ", rank))

					for j := 0; j < 8; j++ {
						idx := offset + j
						if c.wantPos[idx] != board.Pos[idx] {
							file := 'a' + j
							if loc != "" {
								loc += ", "
							}
							loc += fmt.Sprintf("%c%c", file, rank)
						}

						want.WriteString(sq(c.wantPos[idx]))
						got.WriteString(sq(board.Pos[idx]))
					}

					want.WriteByte('\n')
					got.WriteByte('\n')
				}

				var sb strings.Builder
				sb.WriteString(fmt.Sprintf("board differs at %s\nwant:\n%s\ngot:\n%s", loc, want.String(), got.String()))
				t.Error(sb.String())
			}
			if board.ActiveColor != c.wantActiveColor {
				t.Errorf("ActiveColor, want: '%s' got: '%s'", c.wantActiveColor, board.ActiveColor)
			}
			if board.Castling != c.wantCastling {
				t.Errorf("Castling, want: '%s' got: '%s'", c.wantCastling, board.Castling)
			}
			if board.EnPassantSquare != c.wantEnPassantSquare {
				t.Errorf("EnPassantSquare, want: '%s' got: '%s'", c.wantEnPassantSquare, board.EnPassantSquare)
			}
			if board.HalfmoveClock != c.wantHalfMoveClock {
				t.Errorf("HalfmoveClock, want: '%s' got: '%s'", c.wantHalfMoveClock, board.HalfmoveClock)
			}
			if board.FullMove != c.wantFullMove {
				t.Errorf("FullMove, want: '%s' got: '%s'", c.wantFullMove, board.FullMove)
			}
		})
	}
}
