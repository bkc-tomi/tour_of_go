package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n1 := 0
	n2 := 1
	return func() int {
		n3 := n2 + n1
		n0 := n1
		n1 = n2
		n2 = n3
		return n0
	}
}

// FibonacciArray フィボナッチ数列を配列で返す。
func FibonacciArray() func() []int {
	fibonacchi := []int{0, 1}
	return func() []int {
		len := len(fibonacchi)
		n3 := fibonacchi[len-1] + fibonacchi[len-2]
		fibonacchi = append(fibonacchi, n3)
		return fibonacchi
	}
}

func main() {
	f := fibonacci()
	fArray := FibonacciArray()
	for i := 0; i < 50; i++ {
		fmt.Println(f())
		fArray()
	}
	fmt.Println(fArray())
}
