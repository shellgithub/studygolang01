package main

import (
	"encoding/json"
	"fmt"
)

//成员变量名首字母必须大写
type IT struct {
	Comany   string   `json:"-"`        //此字段不会输出到屏幕
	Subjects []string `json:"subjects"` //二次编码
	IsOk     bool     `json:",string"`  //改成字符串
	Price    float64  `json:",string"`
}

func main() {

	jsonBuf := `
	{
	 "subjects": [
	  "go",
	  "C++",
	  "Python",
	  "Test"
	 ],
	 "IsOk": "true",
	 "Price": "6.66"
	}`

	//创建一个map
	m := make(map[string]interface{}, 4)

	err := json.Unmarshal([]byte(jsonBuf), &m)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Printf("m = %+v\n", m)

	//取需要字段
	// var str string
	// str = m["company"]		//err,无法转换
	// fmt.Println("str = ", str)

	fmt.Println("取需要字段")
	var str string
	for key, value := range m {
		// fmt.Println("%v  ======> %v\n", key, value)

		// if key == "company" {
		// 	str = string(value) // cannot convert value (type interface {}) to type string: need type assertion
		// 	fmt.Println("str = ", str)
		// }

		switch data := value.(type) {

		case string:
			str = data
			fmt.Println("map[%s]的值类型为string, value = %s\n", key, str)
		case bool:
			fmt.Println("map[%s]的值类型为string, value = %v\n", key, data)
		case float64:
			fmt.Println("map[%s]的值类型为string, value = %f\n", key, data)
		case []string:
			fmt.Println("map[%s]的值类型为[]string, value = %f\n", key, data)
		case []interface{}:
			fmt.Println("map[%s]的值类型为[]interface{}, value = %f\n", key, data)

		}

	}

}
