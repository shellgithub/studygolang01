package main // 必须

import "fmt"

func test(x int) {
	result := 100 / x
	fmt.Println("result = ", result)
}

// func main() {

// 	fmt.Println("aaaaaaaaaa")
// 	fmt.Println("bbbbbbbbbb")

// 	//调用一个函数，导致内存出问题
// 	test(0)
// 	fmt.Println("cccccccccc")
// }

func main() {

	defer fmt.Println("aaaaaaaaaa")
	defer fmt.Println("bbbbbbbbbb")

	//调用一个函数，导致内存出问题
	defer test(0)

	// 如果一个函数中有多个 defer语句，它们会在 FIFO(后进先出)的顺序执行。
	// 哪怕函数或某个延迟调用发生错误，这些调用依旧会被执行
	defer fmt.Println("cccccccccc")
}
