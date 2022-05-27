package island

import "math/rand"

type heightmap struct {
	Height_map map[vector2]int8
	Radius     int
	Seed       int64
}

func NewHeightMap(radius int, seed int64) *heightmap {
	hm := heightmap{make(map[vector2]int8), radius, seed}
	hm.genHeights()
	return &hm
}

func (hm *heightmap) genHeights() {
	random := rand.New(rand.NewSource(hm.Seed))
	fill := false
	for !fill {
		pos := GetRandomVect(hm.Radius, random)
		val, ok := hm.Height_map[pos]
		if ok {
			if val >= 2 {
				fill = true
			} else {
				hm.Height_map[pos] = val + 1
			}
		} else {
			hm.Height_map[pos] = 1
		}
	}
}
