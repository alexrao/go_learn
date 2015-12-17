package main

import "fmt"

func main() {
	sum := 10
	i := 0

	//for ; sum < 1000; {  // for 语句的这种情况与C的while类似，亦可写成// for sum <1000 {
	// for 死循环
	for {
		i++
		sum += sum
		fmt.Println("Num:", i, "=", sum)

		// if 语句测试
		if sum > 1000 {
			fmt.Println("break i=", i)
			break
		}

	}

	fmt.Printf("The Last Sum:%d\n", sum)
}
