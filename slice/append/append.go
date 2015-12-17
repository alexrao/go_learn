package main

import "fmt"

func main() {
	var a []int
	fmt.Println("a", a)

	// append words to nil slices
	a = append(a, 0)
	printSlice("a", a)

	a = append(a, 2)
	printSlice("a", a)
	fmt.Println(a)

	a = append(a, 10, 20, 30, 400)
	printSlice("a", a)
	fmt.Println(a)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
