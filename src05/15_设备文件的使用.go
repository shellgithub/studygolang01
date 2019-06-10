package main

import (
	"fmt"
	"os"
)

func main() {

	// os.Stdout.Close()        //关闭后无法输出
	// fmt.Println("are u ok?") //往标准输出设备（屏幕）写内容

	//标准设备文件，默认已经打开，用记可以直接使用
	// os.Stdout
	os.Stdout.WriteString(" are u ok?\n")

	// os.Stdin.close()
	var a int
	fmt.Println("请输入a: ")
	fmt.Scan(&a) //从标准输入设备中读取内容，放在a中
	fmt.Println("a = ", a)

}
