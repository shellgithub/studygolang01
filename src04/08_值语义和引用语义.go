package main

import "fmt"

type mystr string //自定义类型，给一个类型改名

type Person struct {
	name string // 名字
	sex  byte   // 性别，字符类型
	age  int    // 年龄
}

//修改成员变量的值

//参数为普通变量，非指针，值语义，一份拷贝
func (p Person) SetInfoValue(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Println("p = ", p)
	fmt.Printf("SetInfoPointer &p = %p\n", &p)
}

//参数为指针变量，引用传递
func (p *Person) SetInfoPointer(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a

	fmt.Printf("SetInfoPointer p = %p\n", p)
}

func main() {
	s1 := Person{"go", 'm', 22}
	fmt.Printf("&s1 = %p\n", &s1) // 打印地址

	//值语义
	// s1.SetInfoValue("mike", 'm', 18)
	// fmt.Println("s1 = ", s1) // 打印内容

	// &s1 = 0xc00000a080
	// p =  {mike 109 18}
	// SetInfoPointer &p = 0xc00000a0a0
	// s1 =  {go 109 22}

	//引用语义
	(&s1).SetInfoValue("mike", 'm', 18)
	fmt.Println("s1 = ", s1) // 打印内容
	// &s1 = 0xc00000a080
	// p =  {mike 109 18}
	// SetInfoPointer &p = 0xc00000a0a0
	// s1 =  {go 109 22}

}
