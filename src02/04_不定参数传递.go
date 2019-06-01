package main // 必须

import "fmt"

func myfunc01(tmp ...int) {
	for _, data := range tmp {
		fmt.Println("data = ", data)
	}

}

func myfunc2(tmp ...int) {
	for _, data := range tmp {
		fmt.Println("data = ", data)
	}
}

func test(args ...int) {
	//全部元素传递给 myfunc
	//myfunc01(args...)

	// 只想把后2 个参数传递给另外一个函数使用
	//myfunc2(args[:2]...) //从args[0-2]开始（不包括2本身）的元素传递过去
	myfunc2(args[2:]...) //从args[2]开始（包括本身），把后面所有元素传递过去
}

func main() {
	test(1, 2, 3, 4)
}
