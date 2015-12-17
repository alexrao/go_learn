package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2} // 类型为 Vertex
	v2 = Vertex{}     // X:0 和 Y:0
	v3 = Vertex{X: 1} // Y:0 被省略
	v4 = Vertex{Y: 1} // X:0 被省略
	v5 = Vertex{Y: 1, X: 3}
	p  = Vertex{1, 2}  // 类型为 *Vertex
	p2 = &Vertex{1, 2} // 类型为 *Vertex
)

func main() {
	fmt.Println(v1, v1.X, v1.Y, p.X, p2.X, v2, v3, v4, v5)
}
