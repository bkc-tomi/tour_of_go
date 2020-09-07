package main

import (
	"fmt"
)

// ErrNegativeSqrt 負の値に関してのerror処理を施した浮動小数点数型
type ErrNegativeSqrt float64

func (err ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", err)
}

// Sqrt 平方根を求める関数。エラー処理あり
func Sqrt(x float64) (float64, error) {
	var z float64 = 1.0
	if x >= 0 {
		for i := 1; i < 10; i++ {
			z -= (z*z - x) / (2 * z)
		}
		return z, nil
	}
	return 0, ErrNegativeSqrt(x)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
