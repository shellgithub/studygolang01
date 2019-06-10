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

	var tmp IT                                   //定义一个结构体变量
	err := json.Unmarshal([]byte(jsonBuf), &tmp) //第二个参数要地址传递
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("tmp = ", tmp)
	fmt.Printf("tmp = %+v\n", tmp)

	//解析需要的字段
	type IT2 struct {
		Subjects []string `json:"subjects`
	}

	var tmp2 IT2
	err = json.Unmarshal([]byte(jsonBuf), &tmp2)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Printf("tmp = %+v\n", tmp2)

}
