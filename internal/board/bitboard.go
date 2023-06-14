package board

import "fmt"

func printBitBoard(bb uint64) {
	var shiftMe uint64 = 1
	rank, file, sq, sq64 := 0, 0, 0, 0

	fmt.Println()
	for rank = rank8; rank >= rank1; rank-- {
		for file = fileA; file <= fileH; file++ {
			sq = fr2sq(file, rank)
			sq64 = sq120toSq64[sq]
			if (shiftMe<<sq64)&bb != 0 {
				fmt.Printf("X")
			} else {
				fmt.Printf("-")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
