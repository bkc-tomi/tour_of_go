package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	fmt.Println("go!")
	c := make(chan int, 30)
	go fibonacci(cap(c), c)

	i := 0
	for n := range c {
		fmt.Println(i, ":", n)
		i++
	}
}
