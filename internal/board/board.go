package board

import (
	"errors"
	"os"

	"github.com/rs/zerolog/log"
)

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

}

func (b *Board) ResetBoard() {
	index := 0
	for index = 0; index < boardSquareNum; index++ {
		b.Pieces[index] = offBoard
	}
	for index = 0; index < 64; index++ {
		b.Pieces[fsq120(index)] = empty
	}
	for index = 0; index < 3; index++ {
		b.BigPieces[index] = 0
		b.MajorPieces[index] = 0
		b.MinorPieces[index] = 0
		b.Pawns[index] = 0
	}
	for index = 0; index < 13; index++ {
		b.PieceNum[index] = 0
	}
	b.KingSquare[white] = noSquare
	b.KingSquare[black] = noSquare

	b.Side = both
	b.EnPassant = noSquare
	b.FiftyMove = 0
	b.Ply = 0
	b.HistPly = 0
	b.CastlePerm = 0
	b.PosKey = 0
}

func (b *Board) GeneratePosKey() uint64 {
	//TODO: https://www.youtube.com/watch?v=WqVwQBXLwE0&list=PLZ1QII7yudbc-Ky058TEaOstZHVbT-2hg&index=12
	return 0
}

func (b *Board) ParseFen(fen string) {

	if fen == "" {
		log.Error().Msg("test")
	}

	rank := rank8
	file := fileA
	piece := 0
	count := 0

	sq64 := 0
	sq120 := 0
	charIndex := 0

	b.ResetBoard()

	// set pieces / blanks for each square
	for charIndex = 0; rank >= rank1; charIndex++ {
		fenChar := fen[charIndex]
		count = 1
		switch fenChar {
		case 'p':
			piece = bP

		case 'r':
			piece = bR

		case 'n':
			piece = bN

		case 'b':
			piece = bB

		case 'k':
			piece = bK

		case 'q':
			piece = bQ

		case 'P':
			piece = wP

		case 'R':
			piece = wR

		case 'N':
			piece = wN

		case 'B':
			piece = wB

		case 'K':
			piece = wK

		case 'Q':
			piece = wQ

		case '1', '2', '3', '4', '5', '6', '7', '8':
			piece = empty
			count = int(fenChar - '0')

		case '/', ' ':
			rank--
			file = fileA
			continue

		default:
			err := errors.New("FEN Error")
			log.Error().Stack().Err(err).Str("fenChar", string(fenChar)).Send()
			os.Exit(1)
		}

		// use count to skip empty squares
		for i := 0; i < count; i++ {
			sq64 = rank*8 + file
			sq120 = fsq120(sq64)
			if piece != empty {
				b.Pieces[sq120] = piece
			}
			file++
		}
	}

	// White or black
	if fen[charIndex] != 'w' && fen[charIndex] != 'b' {
		err := errors.New("FEN Error: Char should be w or b")
		log.Error().Stack().Err(err).Int("charIndex", charIndex).Str("fen[charIndex]", string(fen[charIndex])).Send()
		os.Exit(1)
	}

	if fen[charIndex] == 'w' {
		b.Side = white
	} else {
		b.Side = black
	}

	// Castling permission
	charIndex += 2

	for i := 0; i < 4; i++ {
		if fen[charIndex] == ' ' {
			break
		}
		switch fen[charIndex] {
		case 'K':
			b.CastlePerm |= wKCastling
		case 'Q':
			b.CastlePerm |= wQCastling
		case 'k':
			b.CastlePerm |= bKCastling
		case 'q':
			b.CastlePerm |= bQCastling
		}
		charIndex++
	}
	charIndex++

	if b.CastlePerm < 0 || b.CastlePerm > 15 {
		err := errors.New("FEN Error: CastlePerm invalid")
		log.Error().Stack().Err(err).Int("b.CastlePerm", b.CastlePerm).Send()
		os.Exit(1)
	}

	// en passant square
	if fen[charIndex] != '-' {
		file = int(fen[charIndex] - 'a')
		rank = int(fen[charIndex+1] - '1')

		// TODO: add asserts
		b.EnPassant = fr2sq(file, rank)
	}

	b.PosKey = b.GeneratePosKey()
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

func fsq64(sq120 int) int {
	return sq120toSq64[sq120]
}

func fsq120(sq64 int) int {
	return sq64toSq120[sq64]
}
