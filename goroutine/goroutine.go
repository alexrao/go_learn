package main

import (
	"fmt"
	"time"
)

var count int = 0

func Add(x, y int) {
	if x%2 == 0 {
		time.Sleep(1 * time.Millisecond)
	}
	z := x + y
	fmt.Println(z)
	count++
}

func main() {
	for i := 0; i < 10; i++ {
		go Add(i, i)
	}

	for {
		time.Sleep(100 * time.Millisecond)
		if count >= 9 {
			break
		}
	}
}
