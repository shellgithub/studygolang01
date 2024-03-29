package main // 必须

import "fmt"

func main() {
	// 定义一个数组
	// [数字], 这个数字作为数组元素个数
	var a [10]int
	var b [5]int

	fmt.Printf("len(a) = %d, len(b) = %d\n", len(a), len(b))

	// 注意：定义数组是，指定的数组元素个数必须是常量
	// n := 10
	// var c [n]int //err  non-constant array bound n

	//操作数据元素，从0开始， 到len()-a,不对称元素，这个数字，叫下标
	//这是下标，可以是变量或常量

	//赋值，每个元素
	for i := 0; i < len(a); i++ {
		a[i] = i + 1
	}
	//打印
	//第一个返回下标，第二个返回元素
	for i, data := range a {
		fmt.Printf("a[%d] = %d\n", i, data)
	}

}
