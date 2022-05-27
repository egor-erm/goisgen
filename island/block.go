package island

import "math"

type block struct {
	Id  string
	Pos position
}

type position struct {
	Level int
	X     int
	Y     int
}

type vector2 struct {
	X int
	Y int
}

func (bl *block) getDirection() *direction {
	var vec string
	var zn string

	if math.Abs(float64(bl.Pos.X)) >= math.Abs(float64(bl.Pos.Y)) {
		vec = "x"
		if bl.Pos.X >= 0 {
			zn = "-"
		} else {
			zn = "+"
		}
	} else {
		vec = "y"

		if bl.Pos.Y >= 0 {
			zn = "-"
		} else {
			zn = "+"
		}
	}

	return &direction{vec, zn}
}
