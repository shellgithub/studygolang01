package main

import (
	"fmt"
	"os"
)

func WriteFile(path string) {
	//打开文件，新建文件
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//使用完毕，需要关闭文件
	defer f.Close()

	var buf string

	for i := 0; i < 10; i++ {
		//"i = 1\n", 这个字符串存储在buf中
		buf = fmt.Sprintf("i = %d\n", i)

		// fmt.Println("buf = ", buf)

		n, err := f.WriteString(buf)
		if err != nil {
			fmt.Println("err = ", err)
		}
		fmt.Println("n = ", n)
	}
}

func main() {
	path := "./demo.txt"

	WriteFile(path)

}
