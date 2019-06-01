package main // 必须

import "fmt"

func main() {
	//无参无返回值函数的调用： 函数名()
	MyFunc()
}

func MyFunc() {
	a := 666
	fmt.Println("a = ", a)
}
