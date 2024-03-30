package board

import (
	"fmt"
	"strconv"
)

type Board struct {
	wpawns   uint64
	wbishops uint64
	wknights uint64
	wrooks   uint64
	wqueens  uint64
	wking    uint64

	bpawns   uint64
	bbishops uint64
	bknights uint64
	brooks   uint64
	bqueens  uint64
	bking    uint64

	blacks    [6]*uint64
	whites    [6]*uint64
	allPieces [12]*uint64

	whiteToMove bool

	whiteCastle bool
	blackCastle bool

	enPassantSquare uint8

	reversibleSemiMoves uint8
	fullMoves           uint16
}

func (b *Board) StartPosition() {
	b.wpawns = 0x0000_0000_0000_ff00
	b.wbishops = 0x0000_0000_0000_0024
	b.wknights = 0x0000_0000_0000_0042
	b.wrooks = 0x0000_0000_0000_0081
	b.wqueens = 0x0000_0000_0000_0010
	b.wking = 0x0000_0000_0000_0008

	b.bpawns = 0x00ff_0000_0000_0000
	b.bbishops = 0x2400_0000_0000_0000
	b.bknights = 0x4200_0000_0000_0000
	b.brooks = 0x8100_0000_0000_0000
	b.bqueens = 0x1000_0000_0000_0000
	b.bking = 0x0800_0000_0000_0000
}

func (b *Board) PrintBoard() {
	fmt.Println("Board:")
	var pieces uint64
	for _, p := range b.allPieces {
		pieces ^= *p
	}

	PrintUint64(pieces)
	fmt.Println()
}

func PrintUint64(value uint64) {
	for i := 1; i < 9; i++ {
		binary := int64(value >> (64 - i*8))
		fmt.Printf("%08s\n", strconv.FormatInt(binary&0xff, 2))
	}
}

func NewBoard() *Board {
	b := Board{}

	b.blacks = [6]*uint64{&b.bpawns, &b.bbishops, &b.bknights, &b.brooks, &b.bqueens, &b.bking}
	b.whites = [6]*uint64{&b.wpawns, &b.wbishops, &b.wknights, &b.wrooks, &b.wqueens, &b.wking}
	b.allPieces = [12]*uint64{&b.bpawns, &b.bbishops, &b.bknights, &b.brooks, &b.bqueens, &b.bking,
		&b.wpawns, &b.wbishops, &b.wknights, &b.wrooks, &b.wqueens, &b.wking}
	b.whiteToMove = true
	b.whiteCastle = true
	b.blackCastle = true

	return &b
}
