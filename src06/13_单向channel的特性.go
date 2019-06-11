package main

//go build 13_单向channel的特性.go

func main() {
	//创建一个channel，双向的
	ch := make(chan int)

	//双向chaanel 能隐式转换为单向 channel
	var writeCh chan<- int = ch //只能写，不能读

	writeCh <- 666 //写
	// <-writeCh      //err,  invalid operation: <-writeCh (receive from send-only type chan<- int)

	var readCh <-chan int = ch //只能读，不能写
	<-readCh                   //读
	//readCh <- 666              //写， err, invalid operation: readCh <- 666 (send to receive-only type <-chan int)

	//单向无法转换为双向
	// var ch2 chan int = writeCh //err, cannot use writeCh (type chan<- int) as type chan int in assignment

}
