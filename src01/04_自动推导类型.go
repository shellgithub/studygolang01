package main

import "fmt"

func main() {
	// 赋值，赋值前，必须先声明变量
	var a int
	a = 10 // 赋值，可以使用 N 次
	fmt.Println("a= ", a)

	//:= 自动推导类型，先声明变量b, 再给b 赋值为20
	// 自动推导，同一个变量名只能使用一次，用于初始化那次
	b := 20
	fmt.Println("b= ", b)

	// b := 30  前面已经有变量b, 不能再新建一个变量b
	// 只是赋值，可以的
	b = 30
	fmt.Println("b2 = ", b)

}
