package board

import (
	"fmt"
	"strconv"
)

type Board struct {
	Wpawns   uint64
	Wbishops uint64
	Wknights uint64
	Wrooks   uint64
	Wqueens  uint64
	Wking    uint64

	Bpawns   uint64
	Bbishops uint64
	Bknights uint64
	Brooks   uint64
	Bqueens  uint64
	Bking    uint64

	Blacks    [6]*uint64
	Whites    [6]*uint64
	AllPieces [12]*uint64

	WhiteToMove bool

	WhiteCastle bool
	BlackCastle bool

	EnPassantSquare uint8

	ReversibleSemiMoves uint8
	FullMoves           uint16
}

func (b *Board) StartPosition() {
	b.Wpawns = 0x0000_0000_0000_ff00
	b.Wbishops = 0x0000_0000_0000_0024
	b.Wknights = 0x0000_0000_0000_0042
	b.Wrooks = 0x0000_0000_0000_0081
	b.Wqueens = 0x0000_0000_0000_0010
	b.Wking = 0x0000_0000_0000_0008

	b.Bpawns = 0x00ff_0000_0000_0000
	b.Bbishops = 0x2400_0000_0000_0000
	b.Bknights = 0x4200_0000_0000_0000
	b.Brooks = 0x8100_0000_0000_0000
	b.Bqueens = 0x1000_0000_0000_0000
	b.Bking = 0x0800_0000_0000_0000
}

func (b *Board) PrintBoard() {
	fmt.Println("Board:")
	var pieces uint64
	for _, p := range b.AllPieces {
		pieces ^= *p
	}

	PrintBitboard(pieces)
}

func PrintBitboard(value uint64) {
	for i := 1; i < 9; i++ {
		binary := int64(value >> (64 - i*8))
		fmt.Printf("%08s\n", strconv.FormatInt(binary&0xff, 2))
	}
	fmt.Println()
}

func NewBoard() *Board {
	b := Board{}

	b.Blacks = [6]*uint64{&b.Bpawns, &b.Bbishops, &b.Bknights, &b.Brooks, &b.Bqueens, &b.Bking}
	b.Whites = [6]*uint64{&b.Wpawns, &b.Wbishops, &b.Wknights, &b.Wrooks, &b.Wqueens, &b.Wking}
	b.AllPieces = [12]*uint64{&b.Bpawns, &b.Bbishops, &b.Bknights, &b.Brooks, &b.Bqueens, &b.Bking,
		&b.Wpawns, &b.Wbishops, &b.Wknights, &b.Wrooks, &b.Wqueens, &b.Wking}
	b.WhiteToMove = true
	b.WhiteCastle = true
	b.BlackCastle = true

	return &b
}
