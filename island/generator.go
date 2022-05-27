package island

import (
	"fmt"
	"math/rand"
	"time"
)

type generator struct {
	Island *island
	Hmap   *heightmap
	Seed   int64
}

func (island *island) Generate() {
	island.GenerateIsland(time.Now().UnixNano())
}

func (island *island) GenerateIsland(seed int64) {
	hm := NewHeightMap(island.Radius, seed)

	gen := generator{Island: island, Hmap: hm, Seed: seed}

	gen.fill()
}

func (gen *generator) fill() {
	for key, val := range gen.Hmap.Height_map {
		gen.Island.SetBlock(block{getBlockIdByHeight(val), position{int(val), key.X, key.Y}})
		fmt.Println(key.X, key.Y)
	}
}

func GetRandomInt(radius int, seed int64) int {
	r := rand.New(rand.NewSource(seed))
	a := -radius
	b := radius
	c := a + r.Intn(b-a+1)
	return int(c)
}

func GetRandomVect(radius int, random *rand.Rand) vector2 {
	a := -radius
	b := radius
	c := vector2{a + random.Intn(b-a+1), a + random.Intn(b-a+1)}
	return c
}
