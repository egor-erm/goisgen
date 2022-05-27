package island

import (
	"math"
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
		if gen.isValidPos(key) {
			gen.Island.SetBlock(block{getBlockIdByHeight(val), position{int(val), key.X, key.Y}})
		}
	}
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
