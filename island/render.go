package island

func GetBlockColour(id string) string {
	switch id {
	case "0":
		return "#6EE7FF"
	case "1":
		return "#FFFD96"
	default:
		return "#41FFF6"
	}
}
