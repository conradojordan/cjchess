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
