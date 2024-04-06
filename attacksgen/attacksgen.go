package attacksgen

func GenerateRookAttacks() [64]uint64 {
	var result [64]uint64
	var magicNumber uint64 = 0x0000_0000_0000_0101
	var current uint64 = 0x0101_0101_0101_01fe

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

func GenerateBishopAttacks() [64]uint64 {
	var result [64]uint64
	var magicNumber uint64
	var current uint64 = 0x8040_2010_0804_0200
	var magicNumber2 uint64 = 0x0042_2418_1824_4280

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			result[i*8+j] = current
			magicNumber = 0x0000_0000_0000_0001 << (((i+j)*8)%56 + 8) & 0xffff_ffff_ffff_ffff
			magicNumber ^= 0x0000_0000_0000_0001 << ((64 + 8*i - 8*j) % 72) & 0xffff_ffff_ffff_ffff

			if j == 7 {
				magicNumber2 = (magicNumber2 << 8) & 0xffff_ffff_ffff_ffff
				magicNumber2 = magicNumber2 ^ ((0x0000_0000_0000_0100 >> i) & 0x0000_0000_0000_00ff)
				magicNumber2 = magicNumber2 ^ ((0x0000_0000_0000_0002 << i) & 0x0000_0000_0000_00ff)
				magicNumber = magicNumber2
			}
			current = ((current & 0x7fff_ffff_ffff_ffff) << 1) ^ magicNumber
		}
	}
	return result
}
