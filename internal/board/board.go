package board

const (
	empty = iota
	wP
	wN
	wB
	wR
	wQ
	wK
	bP
	bN
	bB
	bR
	bQ
	bK
)

const (
	fileA = iota
	fileB
	fileC
	fileD
	fileE
	fileF
	fileG
	fileH
	fileNone
)

const (
	rank1 = iota
	rank2
	rank3
	rank4
	rank5
	rank6
	rank7
	rank8
	rankNone
)

const (
	white = iota
	black
	both
)

const (
	a1 = 21 + iota
	b1
	c1
	d1
	e1
	f1
	g1
	h1
)

const (
	a2 = 31 + iota
	b2
	c2
	d2
	e2
	f2
	g2
	h2
)

const (
	a3 = 41 + iota
	b3
	c3
	d3
	e3
	f3
	g3
	h3
)

const (
	a4 = 51 + iota
	b4
	c4
	d4
	e4
	f4
	g4
	h4
)

const (
	a5 = 61 + iota
	b5
	c5
	d5
	e5
	f5
	g5
	h5
)

const (
	a6 = 71 + iota
	b6
	c6
	d6
	e6
	f6
	g6
	h6
)

const (
	a7 = 81 + iota
	b7
	c7
	d7
	e7
	f7
	g7
	h7
)

const (
	a8 = 91 + iota
	b8
	c8
	d8
	e8
	f8
	g8
	h8
	noSquare
)

const (
	WKCastling = 1 << iota
	WQCastling
	BKCastling
	BQCastling
)

type Board struct {
	Pieces     []int
	Pawns      []uint64
	KingSquare []int
	Side       int
	EnPassant  int
	FiftyMove  int
	Ply        int
	HistPly    int

	PosKey      uint64
	PieceNum    []int
	BigPieces   []int
	MajorPieces []int
	MinorPieces []int
}

func initBoard(b *Board) {
	b.Pieces = make([]int, 120)
	b.Pawns = make([]uint64, 3)
	b.KingSquare = make([]int, 2)
	b.PieceNum = make([]int, 13)
	b.BigPieces = make([]int, 3)
	b.MajorPieces = make([]int, 3)
	b.MinorPieces = make([]int, 3)
}
