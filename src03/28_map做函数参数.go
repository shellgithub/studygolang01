package main // 必须

import (
	"fmt"
)

func test(m map[int]string) {
	delete(m, 1)
}

func main() {
	m := map[int]string{1: "mike", 2: "yoyo", 3: "go"}
	fmt.Println("m = ", m)

	test(m)
	fmt.Println("m = ", m)

}
