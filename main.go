package main

import (
	goimage "github.com/egor-erm/goimager/imager"
	"github.com/egor-erm/goisgen/island"
)

func main() {
	img := goimage.NewWithCorners("island.png", -50, -50, 50, 50)
	isl := island.New(50)
	isl.Generate(20)
	for key, val := range isl.Block_list {
		img.SetHexPixel(key.X, key.Y, island.GetBlockColour(val.Id))
	}
	img.Save()
}
