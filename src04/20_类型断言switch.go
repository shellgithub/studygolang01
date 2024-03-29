package main

import "fmt"

type Student struct {
	name string
	id   int
}

func main() {
	//空接口万能类型，保存任意类型的值
	i := make([]interface{}, 3)
	i[0] = 1                    //int
	i[1] = "hello go"           //string
	i[2] = Student{"mike", 666} //Student

	//类型查询，类型断言
	//第一个返回下标，第二个返回下标对应的值，data分别是i[0]， i[1]， i[2]
	for index, data := range i {
		//第一个返回的是值，第二个返回判断结果的真假
		switch value := data.(type) {
		case int:
			fmt.Printf("x[%d]类型为int, 内容为%d\n", index, value)
		case string:
			fmt.Printf("x[%d]类型为string, 内容为%s\n", index, value)
		case Student:
			fmt.Printf("x[%d]类型为Student, 内容为name= %s, id = %d\n", index, value.name, value.id)
		}
	}

}
