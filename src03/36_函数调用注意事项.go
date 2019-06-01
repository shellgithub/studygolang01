package main // 必须

import (
	"fmt"
)

//定义一个结构体类型
type Student struct {
	id   int
	name string
	sex  byte //字符类型
	age  int
	addr string
}

func test02(p *Student) {
	p.id = 666

}

func main() {
	s := Student{1, "mike", 'm', 18, "bj"}

	test02(&s) //地址传递（引用传递），形参无法改实参
	fmt.Println("main: ", s)
}
