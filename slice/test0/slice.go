/**
 * slice 使用
 * slice类似数组，不过不需要指定长度
 * add by rxp -- 20151209
 */

package main

import "fmt"

func main() {
	// 声明数组
	arr := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// slice 声明并赋值
	s := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("S==", s)

	fmt.Println("arr[10]=", arr)

	// s2为slice类型，直接从数组中切片获取，第a[3],a[4]个数据
	s2 := arr[3:5]
	fmt.Println("s2=arr[3:5]==", s2)

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] = %d\n", i, s[i])
	}

	fmt.Println("S==", s)
	fmt.Println("S[1:4]==", s[1:4])
	fmt.Println("S[:3]==", s[:3])
	fmt.Println("S[3:]==", s[3:])

}
