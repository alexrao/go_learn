/**
 * switch 学习测试代码
 * -- by rxp 20151209
 */

package main

import "fmt"

// 带参数的switch
func keys(index int) int {
	var ret int = 0

	switch index {
	case 1:
		ret = 10
		break
	case 2:
		ret = 20
		break
	case 3:
		ret = 30
		break

	default:
		ret = 199
		break

	}
	return ret
}

// 不带参数的switch测试
func swit(index int) int {

	switch {
	case index < 10:
		return 100
	case index < 100:
		return 1000
	default:
		return 0
	}

	return -1
}

func main() {

	var i int

	i = keys(1)
	fmt.Println(i)

	i = swit(9)
	fmt.Println(i)
}
