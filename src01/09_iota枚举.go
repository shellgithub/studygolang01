package main //必须有一个main包

import "fmt" //导入包含，必须要使用

func main() {
	//1、iota 常量自动生成器，每个一行，自动累加1
	//2、iota 给常量赋值使用
	const (
		a = iota //0
		b = iota //1
		c = iota //2
	)
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)

	//3、iota 遇到const，重置为 0
	const d = iota
	fmt.Printf("d = %d\n", d)

	//4、可以只写一个iota
	const (
		a1 = iota //0
		b1
		c1
	)
	fmt.Printf("a1 = %d, b1 = %d, c1 = %d\n", a1, b1, c1)

	//5、如果是同一行，值都一样
	const (
		i          = iota
		j1, j2, j3 = iota, iota, iota
		k          = iota
	)
	fmt.Printf("i = %d, j1 = %d, j2 = %d, j3 = %d, k = %d\n", i, j1, j2, j3, k)

}
