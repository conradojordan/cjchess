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
