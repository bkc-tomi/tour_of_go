package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	vertical := make([][]uint8, dy)

	for y, _ := range vertical {
		horizontal := make([]uint8, dx)
		for x, _ := range horizontal {
			horizontal[x] = uint8(x + y)
		}
		vertical[y] = horizontal
	}

	return vertical
}

func main() {
	pic.Show(Pic)
}
