package main

import (
	"fmt"
)

// Sqrt 平方根を求めて返す関数。エラー処理をしていないので正の数以外は入れない。
func Sqrt(f float64) float64 {
	var z float64 = 1.0
	if f >= 0 {
		for i := 1; i < 10; i++ {
			z -= (z*z - f) / (2 * z)
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(3))
}
