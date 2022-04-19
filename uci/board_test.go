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
			moves: []string{"g1f3"},
			want:  "rnbqkbnr/pppppppp/8/8/8/5N2/PPPPPPPP/RNBQKB1R b KQkq - 1 1",
		},
		{
			start: startPosFEN,
			moves: strings.Split("g1f3 d7d5 e2e3 c7c5 b1c3 g8f6 d2d4 e7e6 f1e2 b8c6", " "),
			want:  "r1bqkb1r/pp3ppp/2n1pn2/2pp4/3P4/2N1PN2/PPP1BPPP/R1BQK2R w KQkq - 2 6",
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
		{
			start: "r1bqkb1r/p4p1p/1p2pnP1/2pp4/1n1P4/2N1PN2/PPP1BPP1/R1BQ1RK1 b kq - 0 9",
			moves: strings.Split("h7h5 g6g7 f6e4 g7h8q", " "),
			want:  "r1bqkb1Q/p4p2/1p2p3/2pp3p/1n1Pn3/2N1PN2/PPP1BPP1/R1BQ1RK1 b q - 0 11",
		},
		{
			start: "r1bqkb1r/p4p1p/1p2pnP1/2pp4/1n1P4/2N1PN2/PPP1BPP1/R1BQ1RK1 b kq - 0 9",
			moves: strings.Split("h7h5 g6g7 f6e4 g7h8q e4g3 h8f8", " "),
			want:  "r1bqkQ2/p4p2/1p2p3/2pp3p/1n1P4/2N1PNn1/PPP1BPP1/R1BQ1RK1 b q - 0 12",
		},
		{
			start: "r1bqkb1r/p4p1p/1p2pnP1/2pp4/1n1P4/2N1PN2/PPP1BPP1/R1BQ1RK1 b kq - 0 9",
			moves: strings.Split("h7h5 g6g7 f6e4 g7h8q e4g3 h8f8 e8d7", " "),
			want:  "r1bq1Q2/p2k1p2/1p2p3/2pp3p/1n1P4/2N1PNn1/PPP1BPP1/R1BQ1RK1 w - - 1 13",
		},
		{
			start: startPosFEN,
			moves: strings.Split("d2d4 g8f6 c2c4 e7e6 g2g3 f8b4 b1d2 d7d5 f1g2 e8g8 g1f3 b7b6 e1g1 c8b7 f3e5 a7a5 d1c2 c7c5 c4d5 b7d5 e2e4 d5b7 d4c5 d8c8 e5d3 b4d2 c1d2 f6e4 a1c1 e4d2 g2b7 c8b7 c2d2 f8d8 d2e3 b6c5 c1c5 b8d7 c5c4 d7b6 c4c5 a8c8 f1c1 h7h6 b2b3 c8c5 d3c5 b7d5 c5e4 b6c8 e4c3 d5a8 e3e4 a8b8 c1d1 c8e7 d1d8 b8d8 g1g2 d8d2 e4f3 e7d5 c3d5 e6d5 h2h4 d2a2 f3d5 g7g6 h4h5 g6h5 d5d8 g8g7 d8d4 g7h7 d4e4 h7h8 e4e8 h8g7 e8e5 g7g6 e5e4 g6g7 e4e5 g7g8 e5h5 a2d2 g2f1 d2c1 f1g2 c1c6 g2g1 c6c1 g1g2 c1d2 h5e5 g8f8 e5h8 f8e7 h8e5 e7d7 g2f3 d2b4 f3g2 b4b7 g2h2 b7f3 e5e1 d7c6 h2g1 f3d5 e1e8 c6c5 e8e1 c5b6 e1e3 b6b5 e3e8 b5b4 e8a4 b4c3 a4a3 d5d1 g1h2 d1h5 h2g2 c3c2 b3b4 h5d5 g2g1 d5d1 g1g2 a5b4 a3b4 d1d5 g2h2 d5h5 h2g2 c2d3 b4b3 d3d4 b3b6 d4c4 b6c6 c4d4 c6a4 d4d3 a4b3 d3e4 f2f3 e4d4 b3a4 d4c3 a4a3 c3c2 a3e3 h5g6 f3f4 g6h5 g2f2 h5d5 g3g4 d5d2 f2f3 d2d7 f4f5 d7d3 f3f4 d3d6 e3e5 d6d2 e5e3 d2d8 e3e4 c2c3 e4e3 c3c4 e3e4 c4c5 e4e5 c5c6 f4e4 d8d6 e5e8 c6c5 e8c8 c5b5 c8c3 f7f6 e4f3 d6d1 f3f4 d1d6 f4e4 d6c6 c3c6 b5c6 e4f4 c6d7 f4g3 d7e8 g3f4 e8f8 f4g3 f8e7 g3h4 e7f7 h4h3 f7e8 h3g3 e8f7 g3h2 f7e7 h2g3 e7e8 g3g2 e8d7 g2h3 d7e7 h3h4 e7f8 h4h3 f8g7 h3g3 h6h5 g4h5 g7h7 g3g4 h7g7 h5h6 g7h6 g4h4 h6h7 h4g3 h7h6 g3h4 h6g7 h4g3 g7h7 g3f4 h7g7 f4g3 g7h8 g3g4 h8h7 g4h3 h7h6 h3g4 h6g7 g4h3 g7f8 h3g4 f8g8 g4h3 g8f7 h3g4 f7f8 g4f4 f8e7 f4e4 e7d6 e4d4 d6c6 d4c4 c6d6 c4d4 d6c6 d4c4 c6d6 c4d4", " "),
			want:  "8/8/3k1p2/5P2/3K4/8/8/8 b - - 39 135",
		},
		{
			start: startPosFEN,
			moves: strings.Split("d2d4 g8f6 c2c4 e7e6 g2g3 f8b4 b1d2 d7d5 f1g2 e8g8 g1f3 b7b6 e1g1 c8b7 f3e5 a7a5 d1c2 c7c5 c4d5 b7d5 e2e4 d5b7 d4c5 d8c8 e5d3 b4d2 c1d2 f6e4 a1c1 e4d2 g2b7 c8b7 c2d2 f8d8 d2e3 b6c5 c1c5 b8d7 c5c4 d7b6 c4c5 a8c8 f1c1 h7h6 b2b3 c8c5 d3c5 b7d5 c5e4 b6c8 e4c3 d5a8 e3e4 a8b8 c1d1 c8e7 d1d8 b8d8 g1g2 d8d2 e4f3 e7d5 c3d5 e6d5 h2h4 d2a2 f3d5 g7g6 h4h5 g6h5 d5d8 g8g7 d8d4 g7h7 d4e4 h7h8 e4e8 h8g7 e8e5 g7g6 e5e4 g6g7 e4e5 g7g8 e5h5 a2d2 g2f1 d2c1 f1g2 c1c6 g2g1 c6c1 g1g2 c1d2 h5e5 g8f8 e5h8 f8e7 h8e5 e7d7 g2f3 d2b4 f3g2 b4b7 g2h2 b7f3 e5e1 d7c6 h2g1 f3d5 e1e8 c6c5 e8e1 c5b6 e1e3 b6b5 e3e8 b5b4 e8a4 b4c3 a4a3 d5d1 g1h2 d1h5 h2g2 c3c2 b3b4 h5d5 g2g1 d5d1 g1g2 a5b4 a3b4", " "),
			want:  "8/5p2/7p/8/1Q6/6P1/2k2PK1/3q4 b - - 0 67",
		},
		{
			start: startPosFEN,
			moves: strings.Split("d2d4 g8f6 c2c4 e7e6 g2g3 f8b4 b1d2 d7d5 f1g2 e8g8 g1f3 b7b6 e1g1 c8b7 f3e5 a7a5 d1c2 c7c5 c4d5 b7d5 e2e4 d5b7 d4c5 d8c8 e5d3 b4d2 c1d2 f6e4 a1c1 e4d2 g2b7 c8b7 c2d2 f8d8 d2e3 b6c5 c1c5 b8d7 c5c4 d7b6 c4c5 a8c8 f1c1 h7h6 b2b3 c8c5 d3c5 b7d5 c5e4 b6c8 e4c3 d5a8 e3e4 a8b8 c1d1 c8e7 d1d8 b8d8", " "),
			want:  "3q2k1/4npp1/4p2p/p7/4Q3/1PN3P1/P4P1P/6K1 w - - 0 30",
		},
		{
			start: startPosFEN,
			moves: strings.Split("d2d4 g8f6 c2c4 e7e6 g2g3 f8b4 b1d2 d7d5 f1g2", " "),
			want:  "rnbqk2r/ppp2ppp/4pn2/3p4/1bPP4/6P1/PP1NPPBP/R1BQK1NR b KQkq - 1 5",
		},
		{
			start: startPosFEN,
			moves: strings.Split("d2d4 g8f6 c2c4 e7e6 g2g3 f8b4 b1d2 d7d5 f1g2 e8g8", " "),
			want:  "rnbq1rk1/ppp2ppp/4pn2/3p4/1bPP4/6P1/PP1NPPBP/R1BQK1NR w KQ - 2 6",
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
