package main // 必须

import "fmt"

func main() {
	// defer 延迟调用，main 函数结束前调用
	defer fmt.Println("bbbbbb")

	fmt.Println("aaaaaaaaaa")

}
