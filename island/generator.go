package island

import (
	"fmt"
	"math/rand"
	"time"
)

type generator struct {
	Island *island
	opora  []block
	random *rand.Rand
}

type direction struct {
	dir  string
	oper string
}

func (island *island) Generate(points int) {
	island.GenerateIsland(points, time.Now().UnixNano())
}

func (island *island) GenerateIsland(points int, seed int64) {
	randomizer := rand.New(rand.NewSource(seed))

	gen := generator{Island: island, opora: []block{}, random: randomizer}

	for i := 0; i < points; i++ {
		x := island.getRandomInt(randomizer)
		y := island.getRandomInt(randomizer)
		fmt.Println(x, y)
		gen.opora = append(gen.opora, block{"1", position{1, x, y}})
	}

	gen.fill()
}

func (gen *generator) fill() {
	for _, bl := range gen.opora {
		vec := bl.getDirection()
		gen.fillBlock(&bl, vec)
	}
}

func (gen *generator) fillBlock(bl *block, vec *direction) {
	var repeat int

	if vec.dir == "x" {
		repeat = bl.Pos.X
	} else {
		repeat = bl.Pos.Y
	}

	for i := 1; i < repeat; i++ {
		if vec.dir == "x" {
			if vec.oper == "+" {
				gen.Island.SetBlock(block{"1", position{1, bl.Pos.X + i, bl.Pos.Y}})
			} else {
				gen.Island.SetBlock(block{"1", position{1, bl.Pos.X - i, bl.Pos.Y}})
			}
		} else {
			if vec.oper == "+" {
				gen.Island.SetBlock(block{"1", position{1, bl.Pos.X, bl.Pos.Y + i}})
			} else {
				gen.Island.SetBlock(block{"1", position{1, bl.Pos.X, bl.Pos.Y - i}})
			}
		}
	}
}

func (island *island) getRandomInt(random *rand.Rand) int {
	a := -island.Radius + 20
	b := island.Radius - 20
	return a + random.Intn(b-a+1)
}
