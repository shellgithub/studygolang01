package main // 必须

import "fmt"

func funcc(c int) {
	fmt.Println("funcb c = ", c)
}

func funcb(b int) {
	funcc(b - 1)
	fmt.Println("funcb b = ", b)
}

func funca(a int) {

	funcb(a - 1)
	fmt.Println("funca a = ", a)

}

func main() {

	funca(3)
	fmt.Println("main")
}
