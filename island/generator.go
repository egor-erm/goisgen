package island

import (
	"math"
	"math/rand"
	"time"
)

type generator struct {
	island *island
	opora  []block
	random *rand.Rand
}

func (island *island) Generate(points int) {
	island.GenerateIsland(points, time.Now().UnixNano())
}

func (island *island) GenerateIsland(points int, seed int64) {
	randomizer := rand.New(rand.NewSource(seed))

	gen := generator{island: island, opora: []block{}, random: randomizer}

	for i := 0; i < points; i++ {
		gen.opora = append(gen.opora, block{"1", position{1, island.getRandomInt(randomizer), island.getRandomInt(randomizer)}})
	}

	gen.fill()
}

func (gen *generator) fill() {

}

func (island *island) getRandomInt(r *rand.Rand) int {
	return r.Intn(island.GetSize()) - int(math.Floor(float64(island.GetSize())/2))
}
