package island

type island struct {
	Biome      string
	Block_list map[vector2]block
	Radius     int
}

func New(radius int) *island {
	return &island{Biome: "", Block_list: make(map[vector2]block), Radius: radius}
}

func (island *island) IsEmpty(x, y int) bool {
	return island.Block_list[vector2{x, y}].Pos.Level == 0
}

func (island *island) GetSize() int {
	return island.Radius
}

func (island *island) SetBlock(bl block) {
	island.Block_list[vector2{bl.Pos.X, bl.Pos.Y}] = bl
}
