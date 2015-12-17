/**
 * slice的range使用
 * range的掉用返回两个参数分别是 (下标,数据 := range slice)
 * add by rxp -- 20151209
 */
package main

import "fmt"

func main() {

	//	var sli = []int{1, 2, 3, 4, 5, 6, 7}
	//
	//	fmt.Println("sli = ", sli)
	//
	//	for i, v := range sli {
	//		fmt.Printf("sli[%d] = %d\n", i, v)
	//	}

	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)

		fmt.Printf("sli[%d] = %d\n", i, pow[i])
	}

}

// _range_ iterates over of elements in a variety of
// data structures. Let's see how to use `range` with some
// of the data structures we've already learned.

//func main() {
//
//	// Here we use `range` to sum the numbers in a slice.
//	// Arrays work like this too.
//	nums := []int{2, 3, 4}
//	sum := 0
//	for _, num := range nums {
//		sum += num
//	}
//	fmt.Println("sum:", sum)
//
//	// `range` on arrays and slices provides both the
//	// index and value for each entry. Above we didn't
//	// need the index, so we ignored it with the
//	// _blank identifier_ `_`. Sometimes we actually want
//	// the indexes though.
//	for i, num := range nums {
//		fmt.Println("index:", i, "num=", num)
//	}
//
//	// `range` on map iterates over key/value pairs.
//	kvs := map[string]string{"a": "apple", "b": "bannana"}
//	for k, v := range kvs {
//		fmt.Printf("%s -> %s\n", k, v)
//	}
//
//	// `range` on strings iterates over Unicode code
//	// points. The first value is the starting byte index
//	// of the `rune` and the second the `rune` itself.
//	for i, c := range "go" {
//		fmt.Println(i, c)
//	}
//}
