package main

import (
	"os"

	"github.com/evandejesus/slimthicc/internal/board"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var chessboard board.Board

func init() {
	chessboard.InitBoard()
}

func main() {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if os.Getenv("GO_ENV") == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// log.Info().Msg("hello world")
	// log.Print(chessboard.BigPieces)

	// f, err := os.Create("example.svg")
	// if err != nil {
	// 	log.Error().Err(err).Send()
	// }
	// defer f.Close()
	// fenStr := "8/5k2/3p4/1p1Pp2p/pP2Pp1P/P4P1K/8/8 b - - 99 50"
	// pos := &chess.Position{}
	// pos.UnmarshalText([]byte(fenStr))
	// image.SVG(f, pos.Board())
}
