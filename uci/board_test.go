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
				writeBoth := func(s string) {
					want.WriteString(s)
					got.WriteString(s)
				}

				writeBoth("   abcdefgh\n   --------\n")
				for i := 0; i < 8; i++ {
					offset := i * 8
					rank := '8' - i
					writeBoth(fmt.Sprintf("%c: ", rank))

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

					writeBoth("\n")
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

func TestMakeMoves(t *testing.T) {
	// arrange
	cases := []struct {
		start string
		moves []string
		want  string
	}{
		{
			start: startPosFEN,
			moves: []string{"g1f3"},
			want:  "rnbqkbnr/pppppppp/8/8/8/5N2/PPPPPPPP/RNBQKB1R b KQkq - 1 1",
		},
		{
			start: startPosFEN,
			moves: strings.Split("g1f3 d7d5 e2e3 c7c5 b1c3 g8f6 d2d4 e7e6 f1e2 b8c6 e1g1", " "),
			want:  "r1bqkb1r/pp3ppp/2n1pn2/2pp4/3P4/2N1PN2/PPP1BPPP/R1BQ1RK1 b kq - 3 6",
		},
		{
			start: "r1bqkb1r/pp3ppp/2n1pn2/2pp4/3P4/2N1PN2/PPP1BPPP/R1BQ1RK1 b kq - 3 6",
			moves: strings.Split("c6b4 h2h4 b7b6 h4h5 g7g5", " "),
			want:  "r1bqkb1r/p4p1p/1p2pn2/2pp2pP/1n1P4/2N1PN2/PPP1BPP1/R1BQ1RK1 w kq g6 0 9",
		},
		{
			start: "r1bqkb1r/p4p1p/1p2pn2/2pp2pP/1n1P4/2N1PN2/PPP1BPP1/R1BQ1RK1 w kq g6 0 9",
			moves: []string{"h5g6"},
			want:  "r1bqkb1r/p4p1p/1p2pnP1/2pp4/1n1P4/2N1PN2/PPP1BPP1/R1BQ1RK1 b kq - 0 9",
		},
	}

	for _, c := range cases {
		t.Run(strings.Join(c.moves, " "), func(t *testing.T) {
			// act
			b := FENtoBoard(c.start)
			b.Moves(c.moves...)
			got := b.FEN()

			// assert
			if c.want != got {
				t.Errorf(fmt.Sprintf("\nwant: '%s'\ngot:  '%s'", c.want, got))
			}
		})
	}
}
