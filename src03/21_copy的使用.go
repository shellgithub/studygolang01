package main // 必须

import (
	"fmt"
)

func main() {
	srcSlice := []int{1, 2}
	dstSlice := []int{6, 6, 6, 6, 6, 6}

	copy(dstSlice, srcSlice)
	fmt.Println("dst = ", dstSlice)

	//打印结果 dst =  [1 2 6 6 6 6]
}
