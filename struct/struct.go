/**
 * 结构提和结构体指针的使用
 * add by rxp 20151209
 */

package main

import "fmt"

// 定义结构体
type STUC struct {
	x int
	y int
}

// 全局使用
var stu STUC

func main() {

	// 结构提函数内部使用
	v := STUC{10, 20}

	stu.x = 11

	// 结构体与指针之间的关系
	p := &stu.y

	*p = 30

	fmt.Println(v.x, v.y)
	fmt.Println("stu:x,y = ", stu.x, stu.y)

}
