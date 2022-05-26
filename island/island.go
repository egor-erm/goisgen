package island

type island struct {
	biome      string
	block_list [][]block
}

func New(size int) *island {
	var px [][]block

	for x := 0; x < size; x++ {
		px = append(px, []block{})
		for y := 0; y < size; y++ {
			px[x] = append(px[x], block{})
		}
	}

	return &island{biome: "", block_list: px}
}

func (island *island) GetSize() int {
	return len(island.block_list)
}

func (island *island) SetBlock(bl block) {
	island.block_list[bl.pos.x][bl.pos.y] = bl
}
