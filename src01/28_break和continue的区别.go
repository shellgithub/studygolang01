package main //必须有一个main包

import "fmt" //导入包含，必须要使用
import "time"

func main() {
	i := 0
	for { //for 后面不写任何东西，这个循环条件永远为真，死循环
		i++
		time.Sleep(time.Second)
		if i == 5 {
			//break //跳出循环，如果嵌套多个循环，跳出最近的那个内循环
			continue //跳出本次循环，下一次继续
		}
		fmt.Println("i = ", i)

	}
}
