package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case <-quit:
			time.Sleep(200 * time.Millisecond)
			fmt.Println("quit")
			return
		case c <- x:
			time.Sleep(200 * time.Millisecond)
			x, y = y, x+y
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	fmt.Println("go start.")

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}
