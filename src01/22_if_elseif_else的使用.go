package main //必须有一个main包

import "fmt" //导入包含，必须要使用

func main() {

	//1
	a := 10
	if a == 10 {
		fmt.Println("a == 10")
	} else {
		fmt.Println("a != 10")
	}

	//2
	if a := 10; a == 10 {
		fmt.Println("a == 10")
	} else {
		fmt.Println("a != 10")
	}

	//3
	a = 18
	if a := 18; a == 10 {
		fmt.Println("a == 10")
	} else if a > 10 {
		fmt.Println("a > 10")
	} else if a < 10 {
		fmt.Println("a < 10")
	}
}
