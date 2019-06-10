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

	//定义一个结构体变量，同时初始化
	s := IT{"itcast", []string{"go", "C++", "Python", "Test"}, true, 6.66}

	//编码，根据内容生成json文本
	//buf, err := json.Marshal(s)
	//格式化打印
	buf, err := json.MarshalIndent(s, "", " ")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	fmt.Println("buf = ", string(buf))
	//将打印的结果使用https://www.json.cn 查看
	/*
		 {"Comany":"itcast","Subjects":["go","C++","Python","Test"],"IsOk":true,"Price":6.66}

		{
		    "Comany":"itcast",
		    "Subjects":[
		        "go",
		        "C++",
		        "Python",
		        "Test"
		    ],
		    "IsOk":true,
		    "Price":6.66
		}
	*/

}
