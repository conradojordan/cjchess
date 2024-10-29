package attacks

import (
	"github.com/conradojordan/cjchess/board"
)

func GenerateRookAttacks() [64]board.Bitboard {
	var result [64]board.Bitboard
	var magicNumber board.Bitboard = 0x101
	var current board.Bitboard = 0x0101_0101_0101_01fe

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			result[i*8+j] = current
			current = (current << 1) ^ magicNumber
		}
		current = (((current ^ 1) << 8) + 1)
		magicNumber = (magicNumber << 8)
	}
	return result
}

func GenerateKnightAttacks() [64]board.Bitboard {
	// - - - - - - - -
	// - a - b - - - -
	// c - - - d - - -
	// - - x - - - - -
	// e - - - f - - -
	// - g - h - - - -
	// - - - - - - - -
	// - - - - - - - -
	var result [64]board.Bitboard
	var current board.Bitboard
	var a, b, c, d, e, f, g, h board.Bitboard

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			current = 0

			a = 1 << (17 + (i*8 + j))
			if j == 7 || i > 5 {
				a = 0
			}
			current ^= a

			b = 1 << (15 + (i*8 + j))
			if j == 0 || i > 5 {
				b = 0
			}
			current ^= b

			c = 1 << (10 + (i*8 + j))
			if j > 5 || i == 7 {
				c = 0
			}
			current ^= c

			d = 1 << (6 + (i*8 + j))
			if j < 2 || i == 7 {
				d = 0
			}
			current ^= d

			e = (1 << 63) >> (69 - (i*8 + j))
			if j > 5 || i == 0 {
				e = 0
			}
			current ^= e

			f = (1 << 63) >> (73 - (i*8 + j))
			if j < 2 || i == 0 {
				f = 0
			}
			current ^= f

			g = (1 << 63) >> (78 - (i*8 + j))
			if i < 2 || j == 7 {
				g = 0
			}
			current ^= g

			h = (1 << 63) >> (80 - (i*8 + j))
			if i < 2 || j == 0 {
				h = 0
			}
			current ^= h

			result[i*8+j] = current
		}
	}
	return result
}

func GenerateBishopAttacks() [64]board.Bitboard {
	var result [64]board.Bitboard
	var magicNumber board.Bitboard
	var current board.Bitboard = 0x8040_2010_0804_0200
	var magicNumber2 board.Bitboard = 0x0042_2418_1824_4280

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			result[i*8+j] = current
			magicNumber = 1 << (((i+j)*8)%56 + 8) & 0xffff_ffff_ffff_ffff
			magicNumber ^= 1 << ((64 + 8*i - 8*j) % 72) & 0xffff_ffff_ffff_ffff

			if j == 7 {
				magicNumber2 = (magicNumber2 << 8) & 0xffff_ffff_ffff_ffff
				magicNumber2 = magicNumber2 ^ ((0x100 >> i) & 0xff)
				magicNumber2 = magicNumber2 ^ ((0x2 << i) & 0xff)
				magicNumber = magicNumber2
			}
			current = ((current & 0x7fff_ffff_ffff_ffff) << 1) ^ magicNumber
		}
	}
	return result
}

func GenerateQueenAttacks() [64]board.Bitboard {
	var result [64]board.Bitboard

	r := GenerateRookAttacks()
	b := GenerateBishopAttacks()

	for i := 0; i < 64; i++ {
		result[i] = r[i] ^ b[i]
	}

	return result
}

func GenerateKingAttacks() [64]board.Bitboard {
	var result [64]board.Bitboard
	var a, b board.Bitboard

	for i := 0; i < 64; i++ {
		a = 7 << (7 + i)
		b = 0xff << ((i/8)*8 + 8)
		result[i] ^= a & b
	}

	for i := 0; i < 63; i++ {
		a = ((5 << 61) >> (62 - i))
		b = 0xff << ((i / 8) * 8)
		result[i] ^= a & b
	}
	result[63] ^= 1 << 62

	for i := 0; i < 64; i++ {
		a = (7 << 54)
		a = a >> (63 - i)
		b = (0xff << ((i / 8) * 8)) >> 8
		result[i] ^= a & b
	}

	return result
}

func GenerateWPawnAttacks() [64]board.Bitboard {
	var current, a, b board.Bitboard
	var results [64]board.Bitboard

	for i := 0; i < 64; i++ {
		current = 0

		a = 1 << (9 + i)
		if i%8 == 7 || i > 56 {
			a = 0
		}
		current ^= a

		b = 1 << (7 + i)
		if i%8 == 0 || i > 56 {
			b = 0
		}
		current ^= b

		results[i] = current
	}

	return results
}

func GenerateBPawnAttacks() [64]board.Bitboard {
	var current, a, b board.Bitboard
	var results [64]board.Bitboard

	for i := 0; i < 64; i++ {
		current = 0

		a = (1 << 63) >> (70 - i)
		if i < 8 || i%8 == 7 {
			a = 0
		}
		current ^= a

		b = (1 << 63) >> (72 - i)
		if i < 8 || i%8 == 0 {
			b = 0
		}
		current ^= b

		results[i] = current
	}

	return results
}

func GenerateWPawnMoves() [64]board.Bitboard {
	var result [64]board.Bitboard

	for i := 0; i < 56; i++ {
		result[i] = 1 << (8 + i)
	}

	for i := 8; i < 16; i++ {
		result[i] = result[i] ^ (1 << (16 + i))
	}

	return result
}

func GenerateBPawnMoves() [64]board.Bitboard {
	var result [64]board.Bitboard

	for i := 8; i < 64; i++ {
		result[i] = 1 << (i - 8)
	}

	for i := 48; i < 56; i++ {
		result[i] = result[i] ^ (1 << (i - 16))
	}

	return result
}
