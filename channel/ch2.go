package main

import (
	"fmt"
)

func task(c chan int) {
	c <- 13
}

func main() {
	ch := make(chan int)
	go task(ch)
	fmt.Println(<-ch)
}
