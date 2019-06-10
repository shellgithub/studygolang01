package main

import "fmt"

func testa() {
	fmt.Println("aaaaaaaaaaaaaa")
}

func testb(x int) {
	// //设置 recover
	defer func() {
		// recover() //可以打印 panic 的错误信息
		// fmt.Println(recover())
		if err := recover(); err != nil { //产生了panic异常
			fmt.Println(err)
		}
	}() //别忘了(), 调用此匿名函数

	var a [10]int
	a[x] = 111 // 当x 为20时候，导致数组越界，产生一个 panic，导致程序崩溃

}

func testc() {
	fmt.Println("ccccccccccccccc")
}

func main() {
	testa()
	testb(10)
	testc()
}
