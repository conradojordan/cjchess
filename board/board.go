package board

import (
	"fmt"
	"strconv"
	"strings"
)

type Bitboard uint64

func (b Bitboard) String() string {
	str := strconv.FormatUint(uint64(b), 2)
	padded_str := fmt.Sprintf("%064s", str)

	var result strings.Builder

	for i := 0; i < 8; i++ {
		result.WriteString(padded_str[i*8 : i*8+8])
		result.WriteString("\n")
	}
	return result.String()
}

type Board struct {
	Wpawns   Bitboard
	Wbishops Bitboard
	Wknights Bitboard
	Wrooks   Bitboard
	Wqueens  Bitboard
	Wking    Bitboard

	Bpawns   Bitboard
	Bbishops Bitboard
	Bknights Bitboard
	Brooks   Bitboard
	Bqueens  Bitboard
	Bking    Bitboard

	Blacks    [6]*Bitboard
	Whites    [6]*Bitboard
	AllPieces [12]*Bitboard

	WhiteToMove bool

	WhiteCastleQueen bool
	WhiteCastleKing  bool
	BlackCastleQueen bool
	BlackCastleKing  bool

	EnPassantSquare Bitboard

	ReversibleSemiMoves uint8
	FullMoves           uint16
}

func (b Board) String() string {
	var result strings.Builder
	result.WriteString("   ABCDEFGH  \n")

	for i := 0; i < 64; i++ {
		current_bit := Bitboard(1 << (63 - i))
		if i%8 == 0 {
			result.WriteByte('\n')
			result.WriteByte(byte((63-i)/8 + 49))
			result.WriteString("  ")
		}
		switch {
		// White pieces
		case (b.Wpawns & current_bit) != 0:
			result.WriteRune('P')
		case (b.Wbishops & current_bit) != 0:
			result.WriteRune('B')
		case (b.Wknights & current_bit) != 0:
			result.WriteRune('N')
		case (b.Wrooks & current_bit) != 0:
			result.WriteRune('R')
		case (b.Wqueens & current_bit) != 0:
			result.WriteRune('Q')
		case (b.Wking & current_bit) != 0:
			result.WriteRune('K')
		// Black pieces
		case (b.Bpawns & current_bit) != 0:
			result.WriteRune('p')
		case (b.Bbishops & current_bit) != 0:
			result.WriteRune('b')
		case (b.Bknights & current_bit) != 0:
			result.WriteRune('n')
		case (b.Brooks & current_bit) != 0:
			result.WriteRune('r')
		case (b.Bqueens & current_bit) != 0:
			result.WriteRune('q')
		case (b.Bking & current_bit) != 0:
			result.WriteRune('k')
		default:
			result.WriteRune('.')
		}

		if i%8 == 7 {
			result.WriteString("  ")
			result.WriteByte(byte((63-i)/8 + 49))
		}
	}
	result.WriteString("\n\n   ABCDEFGH  ")
	return result.String()
}

func (b *Board) InitialPosition() {
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

	b.WhiteToMove = true

	b.WhiteCastleQueen = true
	b.WhiteCastleKing = true
	b.BlackCastleQueen = true
	b.BlackCastleKing = true

	b.EnPassantSquare = 0x0
	b.ReversibleSemiMoves = 0
	b.FullMoves = 0
}

func NewBoard() *Board {
	b := Board{}

	b.Blacks = [6]*Bitboard{&b.Bpawns, &b.Bbishops, &b.Bknights, &b.Brooks, &b.Bqueens, &b.Bking}
	b.Whites = [6]*Bitboard{&b.Wpawns, &b.Wbishops, &b.Wknights, &b.Wrooks, &b.Wqueens, &b.Wking}
	b.AllPieces = [12]*Bitboard{
		&b.Bpawns, &b.Bbishops, &b.Bknights,
		&b.Brooks, &b.Bqueens, &b.Bking,
		&b.Wpawns, &b.Wbishops, &b.Wknights,
		&b.Wrooks, &b.Wqueens, &b.Wking,
	}

	b.InitialPosition()

	return &b
}
