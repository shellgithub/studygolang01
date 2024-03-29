package main

import "fmt"

// go函数可以返回多个值

func test() (a, b, c int) {
	return 1, 2, 3
}

func main() {
	// a := 10
	// b := 20
	// c := 30

	a, b := 10, 20

	// 交换2个变量的值
	var tmp int
	tmp = a
	a = b
	b = tmp
	fmt.Printf("a = %d, b = %d\n", a, b)

	// 上面传统方法

	i, j := 10, 20
	i, j = j, i
	fmt.Printf("i = %d, j = %d\n", i, j)

	//匿名变量，丢弃数据不处理，_ 匿名变量配合函数返回值使用，才有优势
	tmp, _ = i, j
	fmt.Println("tmp = ", tmp)

	var c, d, e int
	c, d, e = test() // return 1,2,3
	fmt.Printf("c = %d, d = %d, e = %d\n", c, d, e)

	//留 d 这个变量
	_, d, _ = test()
	fmt.Printf("d = %d\n", d)

	//留 d e 两个变量
	_, d, e = test()
	fmt.Printf("d = %d, e = %d\n", d, e)

}
