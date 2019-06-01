package main

import "fmt"

type Person struct {
	name string // 名字
	sex  byte   // 性别，字符类型
	age  int    // 年龄
}

func (p Person) SetInfoValue() {
	fmt.Printf("SetInfoPointer")
}

func (p *Person) SetInfoPointer() {
	fmt.Printf("SetInfoPointer")
}

func main() {
	p := Person{"mike", 'm', 18}
	p.SetInfoPointer() //func (p *Person) SetInfoPointer()
	//内部，把p转换成&p后再调用，	(&p).SetInfoPointer()

	p.SetInfoValue()

}

//看到了15_方法的继承
