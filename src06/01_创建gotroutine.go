package main

import (
	"fmt"
	"time"
)

func newTask() {
	for {
		fmt.Println("this is a newTask")
		time.Sleep(time.Second) //延时一秒
	}
}

func main() {

	go newTask() //新建一个协程

	for {
		fmt.Println("this is a goroutine")
		time.Sleep(time.Second) //延时一秒
	}

}
