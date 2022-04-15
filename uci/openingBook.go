package uci

import "strings"

func (u *UCI) BookMove() string {
	// trollfish opening book
	if u.fen == startPosFEN {
		// 1. e4 (White, best (gambits) by test)
		return "e2e4"
	}

	if strings.HasPrefix(u.fen, "rnbqkbnr/pppp1ppp/8/4p3/4P3/8/PPPP1PPP/RNBQKBNR w") {
		// 1. e4 e5 2. Qh5 (White, Wayward Queen)
		return "d1h5"
	}

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

	return ""
}
