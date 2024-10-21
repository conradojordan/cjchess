package main

import (
	"fmt"

	"github.com/conradojordan/cjchess/board"
)

func main() {
	b := board.NewBoard()

	fmt.Println(b)
	fmt.Println(b.Wpawns)
	fmt.Println(b.Brooks)
}
