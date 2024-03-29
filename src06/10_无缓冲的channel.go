package main

import (
	"fmt"
	"time"
	// "runtime"
)

func main() {

	//创建一个无缓存的channel
	ch := make(chan int, 0)

	//len(ch)缓冲剩余数据个数，cap(ch)缓冲区大小
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch))

	//新建协程
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("子协程：i = ", i)
			ch <- i //往 chan 写内容
		}
	}()

	//延时
	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		num := <-ch //读取管理中内容，没有内容前，阻塞
		fmt.Println("num = ", num)
	}

}
