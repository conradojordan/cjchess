package main

import (
	"fmt"

	"github.com/conradojordan/cjchess/board"
	"github.com/conradojordan/cjchess/evaluation"
)

func main() {
	b := board.NewBoard()

	fmt.Println(b)
	fmt.Println(b.Wpawns)

	fmt.Println("Current board evaluation is", evaluation.Evaluate(b))
}
