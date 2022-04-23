package uci

import (
	"math/rand"
	"strings"
)

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
	if strings.HasPrefix(u.fen, "r1b1k1nr/pppp1ppp/2n5/4P3/1b6/2N2N2/PqPBPPPP/1R1QKB1R b") { // Bc2 Bb4 Rb1 ... sac!
		return "b2c3"
	}

	if strings.HasPrefix(u.fen, "r1b1kbnr/pppp1ppp/2n5/4P3/1q6/5N2/PPPBPPPP/RN1QKB1R b") {
		// 1. d4 e5 2. dxe5 Nc6 3. Nf3 Qe7 4. (Bg4, Bg5) Qb4+ 5. Bd2 Qxc2 (Black, Englund Gambit)
		return "b4b2"
	}

	return ""
}

func (u *UCI) BookMove(d4 bool) string {
	if !u.gameAgro {
		move := u.CasualBookMove()
		if move != "" {
			return move
		}
	}

	// trollfish opening book
	if u.fen == startPosFEN {
		if d4 {
			// 1. d4
			return "d2d4"
		}
		if !u.gameAgro {
			// 1. e4 (White, best (gambits) by test)
			return "e2e4"
		}

		whiteMove1 := []string{
			"g1f3", // nf3
			"g2g3", // g3
			"c2c4", // c4
			"e2e3", // e3
		}

		n := rand.Intn(len(whiteMove1))
		return whiteMove1[n]

		/*if rand.Intn(2) == 0 {
			return "e2e4"
		} else {
			return "d2d4"
		}*/
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
		return "b8c8" // Reverse Morra: 1. c4 d5 2. cxd5 c6 3. dxc6 Nxc8
	}

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

	// Learned from games
	if strings.HasPrefix(u.fen, "r1bqkb1r/1p1n1ppp/p1n1p3/2ppP3/3P1P2/2N1BN2/PPP1B1PP/R2QK2R b") {
		return "c5d4" // cxd4; was b5 // checked by the engine A LOT
	}

	if strings.HasPrefix(u.fen, "r1b1k2r/3nbppp/1qn1p3/ppppP3/3P1P2/P1N1BN2/1PP1B1PP/R2Q1RK1 b") {
		return "c5d4" // cxd4; was Ba6
	}

	if strings.HasPrefix(u.fen, "8/8/8/Np1b1kP1/3Kp3/8/8/8 b") {
		return "f5e6" // Ke6 (winning); was b4 (drawn)
	}

	/*if strings.HasPrefix(u.fen, "") {
		return ""
	}

	if strings.HasPrefix(u.fen, "") {
		return ""
	}*/

	if strings.HasPrefix(u.fen, "rnbqkb1r/pp3pp1/2p1pn1p/3p4/2PP3B/2N2N2/PP2PPPP/R2QKB1R b KQkq -") {
		return "d5c4" // 6. ... dxc4 (SF15: d=45,cp=0)
	}

	if strings.HasPrefix(u.fen, "rnbqkb1r/pp3pp1/2p1pn1p/3p2B1/2PP4/2N2N2/PP2PPPP/R2QKB1R w KQkq") {
		return "g5f6" // 6. Bxf6 (SF15: d=45,cp=24 d=40,cp=40); was Bh4 (0.00)
	}

	if strings.HasPrefix(u.fen, "rnbqkb1r/pp3ppp/2p1pn2/3p4/2PP4/2N2N2/PP2PPPP/R1BQKB1R w KQkq") {
		return "e2e3" // 5. e3 (SF15: d=45,cp=35 d=40,cp=40); was Bg5 (d=45,cp=19; d=40,cp=10)
	}

	if strings.HasPrefix(u.fen, "rnbqkb1r/ppp2ppp/4pn2/3p4/2PP4/2N5/PP2PPPP/R1BQKBNR w KQkq") {
		return "c4d5" // 4. cxd6 (0.3, Lichess, depth=44); was Nf3 (0.3)
	}

	if strings.HasPrefix(u.fen, "rnbqkbnr/ppp2ppp/4p3/3p4/2PP4/2N5/PP2PPPP/R1BQKBNR b KQkq") {
		return "c7c6" // 3. ... c6 (0.1, Lichess, depth=43); was Nf6 (0.3)
	}

	if strings.HasPrefix(u.fen, "rnbqkbnr/ppp2ppp/4p3/3p4/2PP4/8/PP2PPPP/RNBQKBNR w KQkq") {
		return "b1c3" // 3. Nc3 (0.1, Lichess, depth=45)
	}

	if strings.HasPrefix(u.fen, "rnbqkbnr/ppp1pppp/8/3p4/2PP4/8/PP2PPPP/RNBQKBNR b KQkq") {
		return "b1c3" // 2. ... e6 (0.4? 0.1?, Lichess, depth=42)
	}

	if strings.HasPrefix(u.fen, "rnbqkbnr/ppp1pppp/8/3p4/3P4/8/PPP1PPPP/RNBQKBNR w KQkq") {
		return "c2c4" // 2. c4 (0.4? 0.3? Lichess, depth=38)
	}

	return ""
}
