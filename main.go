package main

import (
	"os"

	"github.com/evandejesus/slimthicc/internal/board"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var chessboard board.Board

func init() {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if os.Getenv("GO_ENV") == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	chessboard.InitBoard()
}

func main() {

	var b board.Board
	startFen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	b.ParseFen(startFen)
}
