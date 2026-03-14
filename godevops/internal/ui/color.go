package ui

func GetColor(percent float64) string {

	if percent < 50 {
		return "green"
	}

	if percent < 80 {
		return "yellow"
	}

	return "red"
}
