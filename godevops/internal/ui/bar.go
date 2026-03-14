package ui

import "fmt"

// CreateBar gera uma barra visual baseada em percentual
func CreateBar(percent float64, width int) string {

	filled := int((percent / 100) * float64(width))

	bar := ""

	for i := 0; i < width; i++ {

		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}

	color := GetColor(percent)

	return fmt.Sprintf("[%s]%s[-]", color, bar)
}
