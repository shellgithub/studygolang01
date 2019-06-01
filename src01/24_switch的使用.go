package main //必须有一个main包

import "fmt" //导入包含，必须要使用

func main() {
	num := 1

	switch num { //switch后面写的是变量本身

	case 1:
		fmt.Println("按下的是1楼")
		//break // go语言保留了 break 关键字，跳出 switch 语言，不写，默认就包含
		//fallthrough //不跳出 switch 语句，后面的无条件执行
	case 2:
		fmt.Println("按下的是2楼")
	case 3:
		fmt.Println("按下的是3楼")
	case 4:
		fmt.Println("按下的是4楼")
	default:
		fmt.Println("按下的是xxx楼")
	}

}
