package main

import (
	"fmt"
)

func task1(v string, ch chan int) {
	fmt.Println(v)
	ch <- 1
}

func task2(v string, ch chan int) {
	fmt.Println(v)
	ch <- 2
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go task2("task2", ch2)
	go task1("task1", ch1)
	fmt.Println(<-ch1)
	fmt.Println(<-ch2)
}
