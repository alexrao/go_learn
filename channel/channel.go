package main

import (
	"fmt"
)

var Cnt int = 0

func Count(ch chan int) {
	Cnt += 1
	fmt.Println("Counting:", Cnt)
	ch <- 1
}

func main() {
	fmt.Println("Channel Test")

	chs := make([]chan int, 10)

	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}

	for _, ch := range chs {
		<-ch
	}
}
