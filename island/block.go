package island

type block struct {
	id  string
	pos position
}

type position struct {
	level int
	x     int
	y     int
}
