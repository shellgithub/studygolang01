package main //必须有一个main包

import "fmt" //导入包含，必须要使用

func main() {
	// 不同类型变量的声明（定义）

	// var a int
	// var b float64

	var (
		a int
		b float64
	)

	// 自动推导类型
	var (
		a = 10
		b = 3.14
	)

	const (
		i int     = 10
		j float64 = 3.14
	)
	fmt.Println("i = ", i)
	fmt.Println("j = ", j)

	a, b = 10, 3.14
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

}
