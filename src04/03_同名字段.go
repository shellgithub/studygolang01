package main

import "fmt"

type Person struct {
	name string // 名字
	sex  byte   // 性别
	age  int    // 年龄
}

type Student struct {
	Person //只有类型，没有名字，这个匿名字段，继承了 Person的成员
	id     int
	addr   string
	name   string //和 Person 同名了
}

func main() {
	//声明（定义一个变量）
	var s Student

	//默认规则（纠结原则），如果能在本作用域找到此成员，就操作此成员
	//					如果没有找到，找到继承的字段
	s.name = "mike" //操作的是Student的name, 还是Person 的name？结论为 Student
	s.sex = 'm'
	s.age = 18
	s.addr = "bj"

	//显示调试
	fmt.Printf("s = %+v\n", s)

}
