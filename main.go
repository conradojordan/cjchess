package main

import (
	"fmt"

	"github.com/conradojordan/cjchess/attacks"
	"github.com/conradojordan/cjchess/board"
	"github.com/conradojordan/cjchess/evaluation"
)

func main() {
	b := board.NewBoard()

	b.InitialPosition()

	fmt.Println(attacks.KnightAttacks[34])

	fmt.Println(b)

	fmt.Println("\nCurrent board evaluation is", evaluation.Evaluate(b))
}
