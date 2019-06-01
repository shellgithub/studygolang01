package main //必须有一个main包

import "fmt" //导入包含，必须要使用

func main() {

	s := "王思聪"

	// if 和{就是条件，条件通常都是关系运算符}
	if s == "王思聪" {
		fmt.Println("左手+右手")
	}

	//if支持 1 个初始化语句，初始化语句和判断条件以分号分隔
	if a := 10; a == 10 { //条件为真，指向{}语句
		fmt.Println("a == 10")
	}
}
