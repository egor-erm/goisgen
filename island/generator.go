package island

import (
	"math"
	"math/rand"
	"time"
)

type generator struct {
	Island    *island
	Block_map map[vector2]block
	Hmap      *heightmap
	Seed      int64
}

func (island *island) Generate() {
	island.GenerateIsland(time.Now().UnixNano())
}

func (island *island) GenerateIsland(seed int64) {
	hm := NewHeightMap(island.Radius, seed)

	gen := generator{Island: island, Block_map: make(map[vector2]block), Hmap: hm, Seed: seed}

	gen.shape()

	gen.fill()
}

func (gen *generator) shape() {
	for key, val := range gen.Hmap.Height_map {
		if gen.isValidPos(key) {
			gen.Block_map[key] = block{getBlockIdByHeight(val), position{int(val), key.X, key.Y}}
		}
	}
}

func (gen *generator) fill() {

}

func (gen *generator) isValidPos(pos vector2) bool {
	random := rand.New(rand.NewSource(gen.Seed))
	newRad := GetRandomInt(random, gen.Island.Radius-int(math.Round(float64(gen.Island.Radius)*0.5)), gen.Island.Radius)
	return newRad >= int(diag(pos))
}

func diag(vec vector2) float64 {
	return math.Sqrt(math.Pow(float64(vec.X), 2) + math.Pow(float64(vec.Y), 2))
}

func GetRandomInt(random *rand.Rand, a int, b int) int {
	c := a + random.Intn(b-a+1)
	return int(c)
}

func GetRandomVect(radius int, random *rand.Rand) vector2 {
	a := -radius
	b := radius
	c := vector2{a + random.Intn(b-a+1), a + random.Intn(b-a+1)}
	return c
}
