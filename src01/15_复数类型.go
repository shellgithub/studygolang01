package main //必须有一个main包

import "fmt" //导入包含，必须要使用

func main() {
	var t complex128 //声明
	t = 2.1 + 3.15i  //赋值
	fmt.Println("t = ", t)

	//自动推导类型
	t2 := 3.3 + 4.4i
	fmt.Printf("t2 type is %T\n", t2)

	//通过内建函数，取实部和虚部
	fmt.Println("real(t2) = ", real(t2), ", imag(t2) = ", imag(t2))

}
