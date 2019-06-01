package main //必须有一个main包

import "fmt" //导入包含，必须要使用

func main() {
	//给int64 起一个别名叫bigint
	type bigint int64

	var a bigint //等价于int64 a
	fmt.Printf("a type is %T", a)

	type (
		long int64
		char byte
	)

	var b long = 11
	var ch char = 'a'
	fmt.Printf("b = %d, ch = %c\n", b, ch)

}
