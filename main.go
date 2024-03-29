package main

import (
	"github.com/conradojordan/cjchess/board"
)

func main() {
	b := board.NewBoard()

	b.StartPosition()
	b.PrintBoard()
}
