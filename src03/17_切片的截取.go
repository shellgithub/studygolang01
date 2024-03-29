package main // 必须

import (
	"fmt"
)

func main() {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// [low:high:max] 取下标从low 开始的元素， len=high-low, cap=max-low
	s1 := array[:] // [0:len(array):len(array)] 不指定容量，容量和长度一样

	fmt.Println("s1 = ", s1)
	fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))

	//操作某个元素，和数组操作方式一样
	data := array[1]
	fmt.Println("data = ", data)

	data2 := array[2]
	fmt.Println("data2 = ", data2)

	s2 := array[3:6:7] // a[3], a[4], a[5]  len = 6-3=3  cap = 7-3=4
	fmt.Printf("len = %d, cap = %d\n", len(s2), cap(s2))

	s3 := array[0:6:7] // 从0开始，取6个元素，容量也是6
	fmt.Printf("len = %d, cap = %d\n", len(s3), cap(s3))

	s4 := array[3:] //从下标为3开始，到结尾
	fmt.Println("s4 = ", s4)
	fmt.Printf("len = %d, cap = %d\n", len(s4), cap(s4))
}
