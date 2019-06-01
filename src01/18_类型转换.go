package main //必须有一个main包

import "fmt" //导入包含，必须要使用

func main() {

	//这种不能转换的类型，叫不兼容类型
	var flag bool
	flag = true
	fmt.Printf("flag = %d\n", flag)
	fmt.Printf("flag = %t\n", flag)

	//布尔类型不能转换为 int
	//fmt.Printf("flag = %t\n", int(flag))

	// 0 就是假， 非0就是真
	//整形也不能转换为 bool
	//flag = bool(1)

	var ch byte
	ch = 'a' //字符类型本质上就是整形
	var t int
	t = int(ch) //类型转换，把ch的值取出来后，转成int再给t的赋值
	fmt.Println("t = ", t)

}
