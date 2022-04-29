package uci

import (
	"math/rand"
	"strings"
)

type firstMove struct {
	uci  string
	freq int

	min int
	max int
}

var firstMoveMap = []*firstMove{
	// stats from Stockfish 15 depths 35-50 of the starting position

	// move  avg dev    avg  geometric mean  stddev
	// e4       3.81  28.81           28.44    4.67
	// d4       5.76  26.44           25.53    6.77
	// Nf3      2.32  25.94           25.78    2.86
	// g3       3.58  14.44           13.55    4.66
	// c4       4.77  14.88           13.69    5.67
	// e3       3.39  12.88           11.29    4.43

	// top tier
	{uci: "e2e4", freq: 29}, // 1. e4   +0.29
	{uci: "d2d4", freq: 26}, // 1. d4   +0.26
	{uci: "g1f3", freq: 26}, // 1. Nf3  +0.26
	// second tier
	{uci: "c2c4", freq: 14}, // 1. c4   +0.14
	{uci: "g2g3", freq: 14}, // 1. g3   +0.14
	{uci: "e2e3", freq: 6},  // 1. e3   +0.11 // NOTE: has a bad score (46.5%) in Caissabase; discounting eval
	{uci: "b2b3", freq: 6},  // 1. b3    0.00 // NOTE: has a decent score (52.3%) in Caissabase; boosting eval
	// equal
	{uci: "a2a3", freq: 1}, // 1. a3    0.00
	{uci: "c2c3", freq: 1}, // 1. c3    0.00
	{uci: "d2d3", freq: 1}, // 1. d3    0.00
	{uci: "h2h3", freq: 1}, // 1. h3    0.00
	{uci: "b1c3", freq: 1}, // 1. Nc3   0.00
	{uci: "a2a4", freq: 1}, // 1. a4   -0.14
	// bad, but of interest. Komodo beat Stockfish with 1. f4 in 2020
	{uci: "f2f4", freq: 3}, // 1. f4   -0.27
	// bad
	{uci: "b2b4", freq: 0}, // 1. b4   -0.29
	{uci: "h2h4", freq: 0}, // 1. h4   -0.37
	// very bad
	{uci: "b1a3", freq: 0}, // 1. Na3  -0.52
	{uci: "g1h3", freq: 0}, // 1. Nh3  -0.67
	{uci: "f2f3", freq: 0}, // 1. f3   -0.73
	{uci: "g2g4", freq: 0}, // 1. g4   -1.37
}

var firstMoveChoices []string

func init() {
	for _, item := range firstMoveMap {
		if item.freq == 0 {
			continue
		}

		for i := 0; i < item.freq; i++ {
			firstMoveChoices = append(firstMoveChoices, item.uci)
		}
	}
}

func getFirstMove() string {
	n := rand.Intn(len(firstMoveChoices))
	return firstMoveChoices[n]
}

func (u *UCI) BookMove() string {
	if !u.gameAgro {
		move := u.CasualBookMove()
		if move != "" {
			return move
		}
	}

	if u.fen == startPosFEN {
		return getFirstMove()
	}

	return ""
}

func (u *UCI) CasualBookMove() string {
	// Wayward Queen
	if strings.HasPrefix(u.fen, "rnbqkbnr/pppp1ppp/8/4p3/4P3/8/PPPP1PPP/RNBQKBNR w") {
		// 1. e4 e5 2. Qh5 (White, Wayward Queen)
		return "d1h5"
	}

	// Englund Gambit
	if strings.HasPrefix(u.fen, "rnbqkbnr/pppppppp/8/8/3P4/8/PPP1PPPP/RNBQKBNR b") {
		// 1. d4 e5 (Black, Englund Gambit)
		return "e7e5"
	}

	if strings.HasPrefix(u.fen, "rnbqkbnr/pppp1ppp/8/4p3/3P4/8/PPP1PPPP/RNBQKBNR w") {
		// 1. d4 e5 2. dxe5 (White, Englund Gambit)
		return "d4e5"
	}

	if strings.HasPrefix(u.fen, "rnbqkbnr/pppp1ppp/8/4P3/8/8/PPP1PPPP/RNBQKBNR b") {
		// 1. d4 e5 2. dxe5 Nc6 (Black, Englund Gambit)
		return "b8c6"
	}

	if strings.HasPrefix(u.fen, "r1bqkbnr/pppp1ppp/2n5/4P3/8/8/PPP1PPPP/RNBQKBNR w") {
		// 1. d4 e5 2. dxe5 Nc6 3. Nf3 (White, Englund Gambit)
		return "g1f3"
	}

	if strings.HasPrefix(u.fen, "r1bqkbnr/pppp1ppp/2n5/4P3/8/5N2/PPP1PPPP/RNBQKB1R b") { // 3. Nf3
		// 1. d4 e5 2. dxe5 Nc6 3. Nf3 Qe7 (Black, Englund Gambit)
		return "d8e7"
	}

	if strings.HasPrefix(u.fen, "r1bqkbnr/pppp1ppp/2n5/4P3/5B2/8/PPP1PPPP/RN1QKBNR b") { // 3. Bf4
		// 1. d4 e5 2. dxe5 Nc6 3. Bf4 Qe7 (Black, Englund Gambit)
		return "d8e7"
	}

	if strings.HasPrefix(u.fen, "r1b1kbnr/ppppqppp/2n5/4P3/8/5N2/PPP1PPPP/RNBQKB1R w") { // 4. Bg5
		// 1. d4 e5 2. dxe5 Nc6 3. Nf3 Qe7 4. Bg5 (White, Englund Gambit)
		return "c1g5"
	}

	if strings.HasPrefix(u.fen, "r1b1kbnr/ppppqppp/2n5/4P1B1/8/5N2/PPP1PPPP/RN1QKB1R b") { // 4. Bg5 Qb4+
		// 1. d4 e5 2. dxe5 Nc6 3. Nf3 Qe7 4. Bg5 Qb4+ (Black, Englund Gambit)
		return "e7b4"
	}

	if strings.HasPrefix(u.fen, "r1b1kbnr/ppppqppp/2n5/4P3/5B2/5N2/PPP1PPPP/RN1QKB1R b") { // (Nf3, Bf4) ... Qb4+
		// 1. d4 e5 2. dxe5 Nc6 3. Nf3 Qe7 4. Bg4 Qb4+ (Black, Englund Gambit)
		return "e7b4"
	}

	if strings.HasPrefix(u.fen, "r1b1kbnr/pppp1ppp/2n5/4P1B1/1q6/2N2N2/PPP1PPPP/R2QKB1R b") { // Bg5 Nc3
		// 1. d4 e5 2. dxe5 Nc6 3. Nf3 Qe7 4. Bg4 Qb4+ 5. Nc3 Qxc2 (Black, Englund Gambit)
		return "b4b2"
	}

	if strings.HasPrefix(u.fen, "r1b1kbnr/pppp1ppp/2n5/4P3/8/2N2N2/PqPBPPPP/R2QKB1R b") { // Bc2 Bb4
		return "f8b4"
	}

	// TODO: play against humans
	/*if strings.HasPrefix(u.fen, "r1b1k1nr/pppp1ppp/2n5/4P3/1b6/2N2N2/PqPBPPPP/1R1QKB1R b") { // Bc2 Bb4 Rb1 ... sac!
		return "b2c3"
	}*/

	if strings.HasPrefix(u.fen, "r1b1kbnr/pppp1ppp/2n5/4P3/1q6/5N2/PPPBPPPP/RN1QKB1R b") {
		// 1. d4 e5 2. dxe5 Nc6 3. Nf3 Qe7 4. (Bg4, Bg5) Qb4+ 5. Bd2 Qxc2 (Black, Englund Gambit)
		return "b4b2"
	}

	// Smith-Morra Gambit
	if strings.HasPrefix(u.fen, "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b") {
		// 1. e4 c5 (Black, Smith-Morra Gambit)
		return "c7c5"
	}

	if strings.HasPrefix(u.fen, "rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w") {
		// 1. e4 c5 2. d4 (White, Smith-Morra Gambit)
		return "d2d4"
	}

	if strings.HasPrefix(u.fen, "rnbqkbnr/pp1ppppp/8/2p5/3PP3/8/PPP2PPP/RNBQKBNR b") {
		// 1. e4 c5 2. d4 cxd4 (Black, Smith-Morra Gambit)
		return "c5d4"
	}

	if strings.HasPrefix(u.fen, "rnbqkbnr/pp1ppppp/8/8/3pP3/8/PPP2PPP/RNBQKBNR w") {
		// 1. e4 c5 2. d4 cxd4 3. c3 (White, Smith-Morra Gambit)
		return "c2c3"
	}

	if strings.HasPrefix(u.fen, "rnbqkbnr/pp1ppppp/8/8/3pP3/2P5/PP3PPP/RNBQKBNR b") {
		// 1. e4 c5 2. d4 cxd4 3. c3 dxc3 (Black, Smith-Morra Gambit)
		return "d4c3"
	}

	if strings.HasPrefix(u.fen, "rnbqkbnr/pp1ppppp/8/8/4P3/2p5/PP3PPP/RNBQKBNR w") {
		// 1. e4 c5 2. d4 cxd4 3. c3 dxc3 4. Nxc3 (White, Smith-Morra Gambit)
		return "b1c3"
	}

	if strings.HasPrefix(u.fen, "rnbqkbnr/pp2pppp/3p4/8/4P3/2N5/PP3PPP/R1BQKBNR w KQkq -") {
		// Smith-Morra: 4. ... d6 5. Bc4
		return "f1c4"
	}

	if strings.HasPrefix(u.fen, "rnbqkbnr/pp2pppp/3p4/8/2B1P3/2N5/PP3PPP/R1BQK1NR b KQkq -") {
		return "b8c6" // Smith-Morra: 4. ... d6 5. Bc4 Nc6
	}

	// TODO: we don't want to play an alternate move, but we want to respond to one
	if strings.HasPrefix(u.fen, "rnbqkbnr/pp2pppp/3p4/8/2B1P3/2N5/PP3PPP/R1BQK1NR b KQkq -") {
		return "e7e6" // Smith-Morra: 4. ... d6 5. Bc4 e6
	}

	// Reverse Morra
	if strings.HasPrefix(u.fen, "rnbqkbnr/pppppppp/8/8/2P5/8/PP1PPPPP/RNBQKBNR b KQkq -") {
		return "d2d4" // Reverse Morra: 1. c4 d4
	}

	if strings.HasPrefix(u.fen, "rnbqkbnr/ppp1pppp/8/3p4/2P5/8/PP1PPPPP/RNBQKBNR w KQkq -") {
		return "c4d5" // Reverse Morra: 1. c4 d5 2. cxd5
	}

	if strings.HasPrefix(u.fen, "rnbqkbnr/ppp1pppp/8/3P4/8/8/PP1PPPPP/RNBQKBNR b KQkq -") {
		return "c7c6" // Reverse Morra: 1. c4 d5 2. cxd5 c6
	}

	if strings.HasPrefix(u.fen, "rnbqkbnr/pp2pppp/2p5/3P4/8/8/PP1PPPPP/RNBQKBNR w KQkq -") {
		return "d5c6" // Reverse Morra: 1. c4 d5 2. cxd5 c6 3. dxc6
	}

	if strings.HasPrefix(u.fen, "rnbqkbnr/pp2pppp/2P5/8/8/8/PP1PPPPP/RNBQKBNR b KQkq -") {
		return "b8c6" // Reverse Morra: 1. c4 d5 2. cxd5 c6 3. dxc6 Nxc6
	}

	if strings.HasPrefix(u.fen, "r1bqkbnr/pp2pppp/2n5/8/8/2N5/PP1PPPPP/R1BQKBNR b KQkq -") {
		// Reverse Morra: 1. c4 d5 2. cxd5 c6 3. dxc6 Nxc6 4. Nc3
		// { White can play Nc3, d3, e3, g3, a3, h3, e4, Nf3 in this position }
		// 4. ... a6 (alternative to e5 or Nf3)
		return "a7a6"
	}

	/*if strings.HasPrefix(u.fen, "r1bqkbnr/1p2pppp/p1n5/8/8/2N5/PP1PPPPP/R1BQKBNR w KQkq -") {
		// Reverse Morra: 1. c4 d5 2. cxd5 c6 3. dxc6 Nxc6 4. Nc3 a6 5. g3
		// { White can play Nf3, d3, g3, f4, e3 in this position }
		return "g2g3"
	}*/

	// d4 Opening
	if strings.HasPrefix(u.fen, "rnbqkbnr/pppppppp/8/8/3P4/8/PPP1PPPP/RNBQKBNR b") {
		return "g8f6" // 1. d4 Nf6
	}

	if strings.HasPrefix(u.fen, "rnbqkb1r/pppppppp/5n2/8/3P4/8/PPP1PPPP/RNBQKBNR w") {
		return "c2c4" // 1. d4 Nf6 2. c4
	}

	if strings.HasPrefix(u.fen, "rnbqkb1r/pppppppp/5n2/8/2PP4/8/PP2PPPP/RNBQKBNR b") {
		return "e7e6" // 1. d4 Nf6 2. c4 e6
	}

	if strings.HasPrefix(u.fen, "rnbqkb1r/pppppppp/5n2/8/3P4/5N2/PPP1PPPP/RNBQKB1R b") {
		return "e7e6" // 1. d4 Nf6 2. Nf3 e6
	}

	if strings.HasPrefix(u.fen, "rnbqkb1r/pppp1ppp/4pn2/8/2PP4/8/PP2PPPP/RNBQKBNR w") {
		return "g2g3" // 1. d4 Nf6 2. c4 e6 3. g3 ( ... Nf3 )
	}

	if strings.HasPrefix(u.fen, "rnbqkb1r/pppp1ppp/4pn2/8/2PP4/5N2/PP2PPPP/RNBQKB1R b") {
		return "b7b6" // 1. d4 Nf6 2. Nf3 e6 3. c4 b6
	}

	if strings.HasPrefix(u.fen, "rnbqk2r/p1pp1ppp/1p2pn2/8/1bPP4/5NP1/PP2PP1P/RNBQKB1R w") {
		return "c1d2" // 1. d4 Nf6 2. Nf3 e6 3. c4 b6 4. g3 Bb4+ 5. Bd2
	}

	if strings.HasPrefix(u.fen, "rnbqk2r/p1pp1ppp/1p2pn2/8/1bPP4/5NP1/PP1BPP1P/RN1QKB1R b") {
		return "b4e7" // 1. d4 Nf6 2. Nf3 e6 3. c4 b6 4. g3 Bb4+ 5. Bd2 Be7
	}

	if strings.HasPrefix(u.fen, "rnbqkb1r/p1pp1ppp/1p2pn2/8/2PP4/5NP1/PP2PP1P/RNBQKB1R b") {
		return "c8a6" // 1. d4 Nf6 2. Nf3 e6 3. c4 b6 4. g3 Ba6
	}

	if strings.HasPrefix(u.fen, "rn1qkb1r/p1pp1ppp/bp2pn2/8/2PP4/1P3NP1/P3PP1P/RNBQKB1R b") {
		return "d7d5" // 1. d4 Nf6 2. Nf3 e6 3. c4 b6 4. g3 Ba6 5. b3 d5
	}

	if strings.HasPrefix(u.fen, "rn1qkb1r/p1p2ppp/bp2pn2/3p4/2PP4/1P3NP1/P3PPBP/RNBQK2R b") {
		return "b8d7" // 1. d4 Nf6 2. Nf3 e6 3. c4 b6 4. g3 Ba6 5. b3 d5 6. Bg2 Nbd7
	}

	return ""
}
