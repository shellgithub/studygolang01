package main

import "fmt"

func testa() {
	fmt.Println("aaaaaaaaaaaaaa")
}

func testb(x int) {
	var a [10]int
	a[x] = 111

}

func testc() {
	fmt.Println("ccccccccccccccc")
}

func main() {
	testa()
	testb()
	testc()
}
