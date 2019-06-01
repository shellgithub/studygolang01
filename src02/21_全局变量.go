package main // 必须

import "fmt"

func test() {
	fmt.Println("test a =", a)
}

var a int

func main() {
	a = 10
	fmt.Println("a = ", a)
	test()
}
