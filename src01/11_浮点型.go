package main //必须有一个main包

import "fmt" //导入包含，必须要使用

func main() {

	var f1 float32
	f1 = 3.14
	fmt.Println("f1 = ", f1)
	//自动推导类型
	f2 := 3.14
	fmt.Printf("f2 type is %T\n", f2) //f2 type is float64

	// float64存储小数比float32更准确
}
