package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	//b := []int{1, 2, 3, 4, 5}
	//a := make([]int, 5)
	printSlice("a", a)

	b := a[1:4]
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:4]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
