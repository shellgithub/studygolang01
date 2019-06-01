package main //必须有一个main包

import "fmt" //导入包含，必须要使用

func main() {
	//变量，程序运行期间，可以改变的量

	//1、 声明格式  var 变量
	//2、只是声明没有初始化的变量，默认值为 0
	//3、同一个 {} 里，声名的变量名是唯一的
	var a int
	fmt.Println("a=", a)

	a = 10
	fmt.Println("a=", a)
}
