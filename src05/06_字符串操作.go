package main

import (
	"fmt"
	"strings"
)

func main() {
	// "hellogo"中是否包含"hello", 包含返回true, 不包含返回false
	fmt.Println(strings.Contains("Hellogo", "Hello"))
	fmt.Println(strings.Contains("Hellogo", "abc"))

	//Joins 组合
	s := []string{"abc", "hello", "mike", "go"}
	buf := strings.Join(s, "x")
	fmt.Println("buf = ", buf)

	// Index,查找子串的位置
	fmt.Println(strings.Index("abcdhello", "hello"))
	fmt.Println(strings.Index("abcdhello", "go")) //不包含子串返回"-1"

	buf = strings.Repeat("go", 3)
	fmt.Println("buf = ", buf) //  "gogogo"

	// Split 以指定的分隔符拆分
	buf = "hello@abc@go@mike"
	s2 := strings.Split(buf, "@")
	fmt.Println("s2 = ", s2)

	//Trim去掉两头的字符
	buf = strings.Trim("          are u ok ?       ", " ") //去掉两头空格
	fmt.Printf("buf = #%s#\n", buf)

	//Field 切片
	s3 := strings.Fields("          are u ok ?       ") //去掉两头空格
	fmt.Printf("buf = #%s#\n", s3)
	for i, data := range s3 {
		fmt.Println(i, ", ", data)
	}

}
