package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println("Print a[0] a[1]:", a[0], a[1])
	fmt.Println("Print a:", a)
}
