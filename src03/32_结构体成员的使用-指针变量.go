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
	//1、指针有合法指向后，才操作成员
	//先定义一个普通结构体变量

	var s Student
	//再定义一个指针变量，保存s的指针
	var p1 *Student
	p1 = &s
	(*p1).name = "mike"
	p1.sex = 'm'
	p1.age = 18
	p1.addr = "bj"
	fmt.Println("p1 = ", p1)

	//通过 new 申请一个结构体
	p2 := new(Student)
	p2.name = "mike"
	p2.sex = 'm'
	p2.age = 18
	p2.addr = "bj"
	fmt.Println("p2 = ", p2)

}
