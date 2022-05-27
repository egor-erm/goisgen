package island

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

func getBlockIdByHeight(height int8) string {
	switch height {
	case 1:
		return "sand"
	case 2:
		return "grass"
	default:
		return "water"
	}
}
