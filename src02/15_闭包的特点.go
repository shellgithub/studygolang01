package main // 必须

import "fmt"

func test02() func() int {

	var x int //没有初始化，值为 0
	return func() int {
		x++
		return x * x //函数调用完毕，x 自动释放
	}
}

func main() {

	//返回值为一个匿名函数，返回一个函数类型，通过f来调用返回的匿名函数,f来闭包函数
	f := test02()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	// 输出
	// 1
	// 4
	// 9
	// 16
	// 25
}

func test01() int {
	//函数被调用时， x 才分配空间，才初始化为 0
	var x int //没有初始化，值为 0
	x++
	return x * x //函数调用完毕，x 自动释放
}

func main01() {
	fmt.Println(test01())
	fmt.Println(test01())
	fmt.Println(test01())
	fmt.Println(test01())
	fmt.Println(test01())

}
