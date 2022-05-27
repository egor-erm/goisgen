package island

func GetBlockColour(id string) string {
	switch id {
	case "water":
		return "#6EE7FF"
	case "sand":
		return "#FFFD96"
	case "grass":
		return "#BBFF6E"
	default:
		return "#41FFF6"
	}
}
