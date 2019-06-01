package main //必须有一个main包

import "fmt" //导入包含，必须要使用

func main() {
	// 变量： 程序运行期间，可以改变的量，变量声明需要 var
	// 常量： 程序运行期间，不可以改变的量，常量声明需要 const

	const a int = 10
	//a = 20 //err ，常量不允许修改
	fmt.Println("a = ", a)

	const b = 10.2
	fmt.Printf("b type is %T\n", b)
	fmt.Println("b = ", b)

}
