package main //必须有一个main包

import "fmt" //导入包含，必须要使用

func main() {
	//1、声明变量，没有初始化，初始值为false
	var a bool
	fmt.Println("a = ", a)

	a = true
	fmt.Println("a = ", a)

	//2.自动推导类型
	var b = false
	fmt.Println("b = ", b)

	c := false
	fmt.Println("c = ", c)

}
