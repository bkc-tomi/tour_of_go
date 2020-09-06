/*
A Tour of Go Exercise: Loops and Functions
if文とfor文を使用して平方根を求めるプログラムを作る。
*/

package main

import (
	"fmt"
)

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
