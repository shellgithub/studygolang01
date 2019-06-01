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
}

func main() {

	var s1 Student = Student{Person{"mike", 'm', 18}, 1, "bj"}

	fmt.Println("s1 = ", s1)

	//自动推导类型
	s2 := Student{Person{"mike", 'm', 18}, 1, "bj"}
	fmt.Println("s2 = ", s2)
	// %+v,显示更详细
	fmt.Println("s2 = %+v\n", s2)

	//指定成员初始化，没有初始化的常用自动赋值为0
	s3 := Student{id: 1}
	fmt.Println("s3 = %+v\n", s3)

	s4 := Student{Person: Person{name: "mike"}, id: 1}
	fmt.Println("s3 = %+v\n", s4)

	// Student 是结构体，如果 不加  Person 会编译不过
	//s5 := Student{"mike", 'm', 18, 1, "bj"}
	// ./01_匿名字段初始化.go:36:16: cannot use "mike" (type string) as type Person in field value
	// ./01_匿名字段初始化.go:36:29: cannot use 18 (type int) as type string in field value
	// ./01_匿名字段初始化.go:36:33: too many values in Student literal

}
