package board

import "testing"

func TestBoardInitialPosition(t *testing.T) {
	var expected_Wpawns Bitboard = 0x0000_0000_0000_ff00
	var expected_Wbishops Bitboard = 0x0000_0000_0000_0024
	var expected_Wknights Bitboard = 0x0000_0000_0000_0042
	var expected_Wrooks Bitboard = 0x0000_0000_0000_0081
	var expected_Wqueens Bitboard = 0x0000_0000_0000_0010
	var expected_Wking Bitboard = 0x0000_0000_0000_0008

	var expected_Bpawns Bitboard = 0x00ff_0000_0000_0000
	var expected_Bbishops Bitboard = 0x2400_0000_0000_0000
	var expected_Bknights Bitboard = 0x4200_0000_0000_0000
	var expected_Brooks Bitboard = 0x8100_0000_0000_0000
	var expected_Bqueens Bitboard = 0x1000_0000_0000_0000
	var expected_Bking Bitboard = 0x0800_0000_0000_0000

	board := NewBoard()
	board.InitialPosition()

	// White pieces
	if board.Wpawns != expected_Wpawns {
		t.Errorf("white pawns are not in correct initial position. Expected %v, got %v", expected_Wpawns, board.Wpawns)
	}
	if board.Wbishops != expected_Wbishops {
		t.Errorf("white bishops are not in correct initial position. Expected %v, got %v", expected_Wbishops, board.Wbishops)
	}
	if board.Wknights != expected_Wknights {
		t.Errorf("white knights are not in correct initial position. Expected %v, got %v", expected_Wknights, board.Wknights)
	}
	if board.Wrooks != expected_Wrooks {
		t.Errorf("white rooks are not in correct initial position. Expected %v, got %v", expected_Wrooks, board.Wrooks)
	}
	if board.Wqueens != expected_Wqueens {
		t.Errorf("white queens are not in correct initial position. Expected %v, got %v", expected_Wqueens, board.Wqueens)
	}
	if board.Wking != expected_Wking {
		t.Errorf("white king are not in correct initial position. Expected %v, got %v", expected_Wking, board.Wking)
	}

	// Black pieces
	if board.Bpawns != expected_Bpawns {
		t.Errorf("black pawns are not in correct initial position. Expected %v, got %v", expected_Bpawns, board.Bpawns)
	}
	if board.Bbishops != expected_Bbishops {
		t.Errorf("black bishops are not in correct initial position. Expected %v, got %v", expected_Bbishops, board.Bbishops)
	}
	if board.Bknights != expected_Bknights {
		t.Errorf("black knights are not in correct initial position. Expected %v, got %v", expected_Bknights, board.Bknights)
	}
	if board.Brooks != expected_Brooks {
		t.Errorf("black rooks are not in correct initial position. Expected %v, got %v", expected_Brooks, board.Brooks)
	}
	if board.Bqueens != expected_Bqueens {
		t.Errorf("black queens are not in correct initial position. Expected %v, got %v", expected_Bqueens, board.Bqueens)
	}
	if board.Bking != expected_Bking {
		t.Errorf("black king are not in correct initial position. Expected %v, got %v", expected_Bking, board.Bking)
	}
}

func TestBoardStringRepresentation(t *testing.T) {
	board := NewBoard()
	board.InitialPosition()

	got := board.String()
	expected := `   ABCDEFGH  

8  rnbqkbnr  8
7  pppppppp  7
6  ........  6
5  ........  5
4  ........  4
3  ........  3
2  PPPPPPPP  2
1  RNBQKBNR  1

   ABCDEFGH  `

	if got != expected {
		t.Errorf("Failed board string representation test. Expected %s, got %s", expected, got)
	}
}
