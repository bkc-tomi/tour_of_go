package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for i := 0; i < 10; i++ {
		time.Sleep(200 * time.Millisecond)
		c <- x
		x, y = y, x+y
	}
	quit <- 0
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	fmt.Println("go start.")

	go fibonacci(c, quit)
	func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case <-quit:
				fmt.Println("quit")
				return
			}
		}
	}()
}
