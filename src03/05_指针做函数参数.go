package main // 必须

import "fmt"

func swap(p1, p2 *int) {
	*p1, *p2 = *p2, *p1

}

func main() {
	a, b := 10, 20

	//通过一个函数交换a和b 的内容
	swap(&a, &b) //变量本身传递，值传递（站在变量角度）
	fmt.Printf("main: a = %d, b = %d\n", a, b)

}
