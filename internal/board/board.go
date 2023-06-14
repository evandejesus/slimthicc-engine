package board

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

var Sq120toSq64 [boardSquareNum]int
var Sq64toSq120 [64]int

const (
	maxGameMoves   = 2048
	boardSquareNum = 120
)

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
	Pieces      [boardSquareNum]int
	Pawns       [3]uint64
	KingSquare  [2]int
	Side        int
	EnPassant   int
	FiftyMove   int
	Ply         int
	HistPly     int
	PosKey      uint64
	PieceNum    [13]int
	BigPieces   [3]int
	MajorPieces [3]int
	MinorPieces [3]int
	CastlePerm  int
	History     [maxGameMoves]Undo

	PieceList [13][10]int
}

type Undo struct {
	Move       int
	CastlePerm int
	EnPassant  int
	FiftyMove  int
	PosKey     int
}

func (b *Board) InitBoard() {
	initSq120toSq64()

	var playBitBoard uint64 = 0
	log.Info().Msg("Start")
	PrintBitBoard(playBitBoard)
	playBitBoard = playBitBoard | (uint64(1) << Sq120toSq64[d2])
	log.Info().Msg("D2 Added")
	PrintBitBoard(playBitBoard)
	playBitBoard = playBitBoard | (uint64(1) << Sq120toSq64[g2])
	log.Info().Msg("G2 Added")
	PrintBitBoard(playBitBoard)
}

func printLists() {
	for i := 0; i < boardSquareNum; i++ {
		if i%10 == 0 {
			fmt.Println()
		}
		fmt.Printf("%5d", Sq120toSq64[i])
	}

	fmt.Println()
	fmt.Println()

	for i := 0; i < 64; i++ {
		if i%8 == 0 {
			fmt.Println()
		}
		fmt.Printf("%5d", Sq64toSq120[i])
	}
}

func initSq120toSq64() {
	var index int
	sq := a1
	sq64 := 0
	for index = 0; index < boardSquareNum; index++ {
		Sq120toSq64[index] = 64
	}
	for index = 0; index < 64; index++ {
		Sq64toSq120[index] = 120
	}

	for rank := rank1; rank <= rank8; rank++ {
		for file := fileA; file <= fileH; file++ {
			sq = fr2sq(file, rank)
			Sq64toSq120[sq64] = sq
			Sq120toSq64[sq] = sq64
			sq64 += 1
		}
	}
}

func fr2sq(f, r int) int {
	return 21 + f + r*10
}

func sq64(sq120 int) int {
	return Sq120toSq64[sq120]

}
