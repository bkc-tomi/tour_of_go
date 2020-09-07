package main

import (
	"golang.org/x/tour/pic"
)

// Pic 縦横の要素数を受け取り２次元配列を返す関数。各要素の値は自由に決めて良い
func Pic(dx, dy int) [][]uint8 {
	vertical := make([][]uint8, dy)
	for y, _ := range vertical {
		horizontal := make([]uint8, dx)
		for x, _ := range horizontal {
			// 引数は自由に決める。
			horizontal[x] = uint8(x)
		}
		vertical[y] = horizontal
	}

	return vertical
}

func main() {
	pic.Show(Pic)
}
