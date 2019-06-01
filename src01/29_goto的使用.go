package main //必须有一个main包

import "fmt" //导入包含，必须要使用

func main() {

	//goto 可以用在任何地方，但是不能跨函数使用
	fmt.Println("11111111")

	goto End //goto是关键字，End是用户起的名字，叫标签
	fmt.Println("222222222")
End:
	fmt.Println("33333333")
}
