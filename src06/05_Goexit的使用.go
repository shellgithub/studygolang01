package main

import (
	"fmt"
	"runtime"
	// "runtime"
)

func test() {
	defer fmt.Println("ccccccccccc")

	// return //终止此函数
	runtime.Goexit() //终止所在协程

	fmt.Println("dddddddddd")
}

func main() {

	//创建新建的协程
	go func() {

		fmt.Println("aaaaaaaaaaaa")

		test()

		fmt.Println("bbbbbbbb")

	}()

	//特地写一个死循环，目的不让主协程结束

	for {

	}

}
