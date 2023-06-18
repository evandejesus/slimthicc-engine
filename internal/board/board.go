package board

import (
	"fmt"
	"math/rand"

	"github.com/rs/zerolog/log"
)

var sq120toSq64 [boardSquareNum]int
var sq64toSq120 [64]int
var setMask [64]uint64
var clearMask [64]uint64

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
	wKCastling = 1 << iota
	wQCastling
	bKCastling
	bQCastling
)

// Board is board
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

func initBitMasks() {
	var index int = 0
	for index = 0; index < 64; index++ {
		setMask[index] = setMask[index] | (uint64(1) << uint64(index))
		clearMask[index] = ^setMask[index]
	}
}

func (b *Board) InitBoard() {
	initSq120toSq64()
	initBitMasks()

	var playBitBoard uint64 = 0

	// for index := 0; index < 64; index++ {
	// 	log.Info().Int("index", index).Send()
	// 	PrintBitBoard(clearMask[index])
	// }

	var pieceOne int32 = rand.Int31()
	var pieceTwo int32 = rand.Int31()
	var pieceThree int32 = rand.Int31()
	var pieceFour int32 = rand.Int31()
	var key int32 = pieceOne ^ pieceTwo ^ pieceFour
	var tempKey int32 = pieceOne
	tempKey ^= pieceThree
	tempKey ^= pieceThree
	tempKey ^= pieceFour
	tempKey ^= pieceTwo
	log.Info().Msg(fmt.Sprintf("key:%X", key))
	log.Info().Msg(fmt.Sprintf("tempKey:%X", tempKey))

	setBit(&playBitBoard, 61)
	PrintBitBoard(playBitBoard)

}

func setBit(bb *uint64, sq int) {
	*bb = *bb | setMask[sq]
}

func clearBit(bb *uint64, sq int) {
	*bb = *bb & clearMask[sq]
}

func initSq120toSq64() {
	var index int
	sq := a1
	sq64 := 0
	for index = 0; index < boardSquareNum; index++ {
		sq120toSq64[index] = 64
	}
	for index = 0; index < 64; index++ {
		sq64toSq120[index] = 120
	}

	for rank := rank1; rank <= rank8; rank++ {
		for file := fileA; file <= fileH; file++ {
			sq = fr2sq(file, rank)
			sq64toSq120[sq64] = sq
			sq120toSq64[sq] = sq64
			sq64 += 1
		}
	}
}

func fr2sq(f, r int) int {
	return 21 + f + r*10
}

func sq64(sq120 int) int {
	return sq120toSq64[sq120]

}
