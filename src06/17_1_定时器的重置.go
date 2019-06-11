package main

import (
	"fmt"
	"time"
)

func main() {

	timer := time.NewTimer(3 * time.Second)
	ok := timer.Reset(1 * time.Second) //重新设置为1s

	fmt.Println("ok = ", ok)

	<-timer.C
	fmt.Println("时间到")

}
