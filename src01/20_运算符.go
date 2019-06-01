package main //必须有一个main包

import "fmt" //导入包含，必须要使用

func main() {

	fmt.Println("4 > 3 结果：", 4 > 3)
	fmt.Println("4 !=3 ", 4 != 3)

	fmt.Println("!(4 > 3) 结果：", !(4 > 3))
	fmt.Println("!(4 !=3) ", !(4 != 3))

	//&& 与，并且，左边右边都为真，结果才为真，其它都为假
	fmt.Println("true && true 结果：", true && true)
	fmt.Println("true && false 结果：", true && false)

	//  ||, 或者，左边右边都为假，结果才为假，其它都为真
	fmt.Println("true || true 结果：", true || true)
	fmt.Println("true || false 结果：", true || false)

	a := 8
	//fmt.Println("0 <= a <= 10  结果为：", 0 <= a <= 10)
	//./20_运算符.go:22:51: cannot convert 10 (type untyped number) to type bool
	//./20_运算符.go:22:51: invalid operation: a <= a <= 10 (mismatched types bool and int)

	fmt.Println("0 <= a <= 10  结果为：", 0 <= a && a <= 10)

}
