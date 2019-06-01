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

func main() {

	//顺序初始化，每个成员必须初始化,别忘了 &
	var p1 *Student = &Student{1, "mike", 'm', 18, "bj"}
	fmt.Println("*p1 = ", *p1)

	//指定成员初始化，没有初始化的成员，自动赋值为0
	var p2 = &Student{name: "mike", addr: "bj"}
	fmt.Printf("p2 type is %T\n", p2)
	fmt.Println("*p2 = ", *p2)

	// *p1 =  {1 mike 109 18 bj}
	// p2 type is *main.Student
	// *p2 =  {0 mike 0 0 bj}
}
