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

	whiteToMove     bool
	whiteCastle     bool
	blackCastle     bool
	numMoves        uint16
	reversibleMoves uint8
}

func (b *Board) StartPosition() {
	b.wpawns = 0x000000000000ff00
	b.wbishops = 0x0000000000000024
	b.wknights = 0x0000000000000042
	b.wrooks = 0x0000000000000081
	b.wqueens = 0x0000000000000010
	b.wking = 0x0000000000000008

	b.bpawns = 0x00ff000000000000
	b.bbishops = 0x2400000000000000
	b.bknights = 0x4200000000000000
	b.brooks = 0x8100000000000000
	b.bqueens = 0x1000000000000000
	b.bking = 0x0800000000000000
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
