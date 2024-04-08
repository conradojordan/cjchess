package main

import (
	"fmt"

	"github.com/conradojordan/cjchess/attacksgen"
	"github.com/conradojordan/cjchess/board"
)

func main() {
	b := board.NewBoard()

	b.StartPosition()
	b.PrintBoard()

	// board.PrintBitboard(b.Wpawns)
	// board.PrintBitboard(b.Brooks)

	a := attacksgen.GenerateBPawnAttacks()

	// for _, v := range a {
	// 	board.PrintBitboard(v)
	// 	fmt.Scanln()
	// }

	fmt.Println("var BPawnAttacks [64]uint64 = [64]uint64{")
	for i := 0; i < 64; i += 4 {
		fmt.Printf("	0x%016x, 0x%016x, 0x%016x, 0x%016x,\n", a[i], a[i+1], a[i+2], a[i+3])
	}
	fmt.Println("}")
}
