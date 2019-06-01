package main // 必须

import "fmt"

func MyFunc01(a int, b int) { //固定参数

}

// ... int 类型这样的类型， ... type 不定参数类型
func MyFunc02(args ...int) { //传递的实参可以是0或多个
	fmt.Println("len(args) = ", len(args))
	for i := 0; i < len(args); i++ {
		fmt.Printf("args[%d] = %d\n", i, args[i])
	}

	fmt.Println("==============================")

	for i, data := range args {
		fmt.Printf("args[%d] = %d\n", i, data)
	}

}

func main01() {
	//MyFunc01(666, 777)

	MyFunc02()
	fmt.Println("+++++++++++++++++++++++++++++")
	MyFunc02(1)
	MyFunc02(1, 2, 3)

}

// 固定参数一定要传参，不定参数根据需求传递
func MyFunc03(a int, args ...int) {

}

// 注意：不定参数，一定（只能）放在形参中的最后一个参数
// func MyFunc04(args   int, args ...int) {

// }

func main() {
	MyFunc03(1111)
}
