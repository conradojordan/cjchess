package evaluation

import (
	"math/bits"

	"github.com/conradojordan/cjchess/board"
)

func EvaluateMaterial(board *board.Board) float32 {
	whiteScore := (bits.OnesCount64(uint64(board.Wpawns)) +
		3*bits.OnesCount64(uint64(board.Wbishops)) +
		3*bits.OnesCount64(uint64(board.Wknights)) +
		5*bits.OnesCount64(uint64(board.Wrooks)) +
		9*bits.OnesCount64(uint64(board.Wqueens)))

	blackScore := (bits.OnesCount64(uint64(board.Bpawns)) +
		3*bits.OnesCount64(uint64(board.Bbishops)) +
		3*bits.OnesCount64(uint64(board.Bknights)) +
		5*bits.OnesCount64(uint64(board.Brooks)) +
		9*bits.OnesCount64(uint64(board.Bqueens)))

	return float32(whiteScore) - float32(blackScore)
}

func Evaluate(board *board.Board) float32 {
	return EvaluateMaterial(board)
}
