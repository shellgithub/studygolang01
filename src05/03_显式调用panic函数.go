package main

import "fmt"

func testa() {
	fmt.Println("aaaaaaaaaaaaaa")
}

func testb() {
	// fmt.Println("bbbbbbbbbbbbbbb")
	//显示调用panic 函数，导致程序中断
	panic("this is a panic test")
}

func testc() {
	fmt.Println("ccccccccccccccc")
}

func main() {
	testa()
	testb()
	testc()
}
