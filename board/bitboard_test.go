package board

import "testing"

func TestBitboardStringRepresentation(t *testing.T) {
	var test_scenarios = []struct {
		bitboard Bitboard
		expected string
	}{
		{0x0102_0408_1020_4080, `00000001
00000010
00000100
00001000
00010000
00100000
01000000
10000000
`},
	}

	for _, scenario := range test_scenarios {
		if string_representation := scenario.bitboard.String(); string_representation != scenario.expected {
			t.Errorf("got %v, want %v", string_representation, scenario.expected)
		}
	}
}
