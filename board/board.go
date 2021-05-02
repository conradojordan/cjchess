package board

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

	blacks    [6]uint64
	whites    [6]uint64
	allPieces [12]uint64
}

func NewBoard() *Board {
	b := Board{}

	b.blacks = [6]uint64{b.bpawns, b.bbishops, b.bknights, b.brooks, b.bqueens, b.bking}
	b.whites = [6]uint64{b.wpawns, b.wbishops, b.wknights, b.wrooks, b.wqueens, b.wking}
	b.allPieces = [12]uint64{b.bpawns, b.bbishops, b.bknights, b.brooks, b.bqueens, b.bking,
		b.wpawns, b.wbishops, b.wknights, b.wrooks, b.wqueens, b.wking}

	return &b
}
