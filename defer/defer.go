/**
 * defer 调用测试
 * add by rxp --20151209
 */
package main

import "fmt"

func test() int {
	defer fmt.Println("Defer 0 Test")
	defer fmt.Println("Defer 1 Test")

	fmt.Println("No Defer 0 test ")
	fmt.Println("No Defer 1 test ")

	return 1
}

func defer_for() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Defer for function")
}

func main() {

	// test 1
	fmt.Println("Main Print:", test())

	// test 2
	defer_for()
}
