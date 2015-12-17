package main

import "fmt"

func main() {

	i := 32

	p := &i

	fmt.Println("p=", p, ",*p", *p)

	*p = 20
	fmt.Println("p=", p, ",*p", *p, "i=", i, "&i = ", &i)
}
