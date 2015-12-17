package main

import "fmt"

func main() {
	var b int = 15
	var a int

	numbers := [6]int{1, 2, 3, 5}

	// loop
	for a := 0; a < 10; a++ {
		fmt.Printf("1 Value of a:%d\n", a)
	}

	fmt.Printf("1 a = %d, b=%d\n", a, b)

	for a < b {
		a++
		fmt.Printf("2 Value of a:%d, b=%d\n", a, b)
	}

	for i, x := range numbers {
		fmt.Printf("3 Value of x=%d at %d\n", x, i)
	}
}
