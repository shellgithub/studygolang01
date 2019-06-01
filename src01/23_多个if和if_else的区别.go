package main //必须有一个main包

import "fmt" //导入包含，必须要使用

func main() {

	//1
	a := 18
	if a == 10 {
		fmt.Println("a == 10")
	} else if a > 10 {
		fmt.Println("a > 10")
	} else if a < 10 {
		fmt.Println("a < 10")
	}

	//2
	b := 18
	if b := 18; b == 10 {
		fmt.Println("b == 10")
	}

	if b > 10 {
		fmt.Println("b > 10")
	}

	if b < 10 {
		fmt.Println("b < 10")
	}

}
