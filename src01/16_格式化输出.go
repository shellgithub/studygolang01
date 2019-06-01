package main //必须有一个main包

import "fmt" //导入包含，必须要使用

func main() {
	a := 10
	b := "abc"
	c := 'a'
	d := 2.14
	//%T 操作变量所属类型
	fmt.Printf("%T,%T,%T,%T\n", a, b, c, d)

	//%d 整形格式
	//%s 字符串格式
	//%c 字符个数
	//%f 浮点型个数
	fmt.Printf("a = %d, b = %s, c = %c, d = %f\n", a, b, c, d)
	//%v 自动匹配格式输出
	fmt.Printf("a = %v, b = %v, c = %v, d = %v\n", a, b, c, d)

}
