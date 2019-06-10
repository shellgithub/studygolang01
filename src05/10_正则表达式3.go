package main

import (
	"fmt"
	"regexp"
)

func main() {

	//`` 原生字符串
	buf := `<!DOCTYPE html>
<html>
 <head> 
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" /> 
  <title>Go库文档</title> 
  <link type="text/css" rel="stylesheet" href="./index_files/style.css" /> 
  <style type="text/css"></style> 
  <style>@-moz-keyframes nodeInserted{from{opacity:0.99;}to{opacity:1;}}@-webkit-keyframes nodeInserted{from{opacity:0.99;}to{opacity:1;}}@-o-keyframes nodeInserted{from{opacity:0.99;}to{opacity:1;}}@keyframes nodeInserted{from{opacity:0.99;}to{opacity:1;}}embed,object{animation-duration:.001s;-ms-animation-duration:.001s;-moz-animation-duration:.001s;-webkit-animation-duration:.001s;-o-animation-duration:.001s;animation-name:nodeInserted;-ms-animation-name:nodeInserted;-moz-animation-name:nodeInserted;-webkit-animation-name:nodeInserted;-o-animation-name:nodeInserted;}</style>
  <style type="text/css" id="8726384429"></style>
 </head> 
 <body > 
  <div id="lowframe" style="position: fixed; bottom: 0; left: 0; height: 0; width: 100%; border-top: thin solid grey; background-color: white; overflow: auto;">
    ... 
  </div>
  <!-- #lowframe --> 

  <div id="page" class="wide" tabindex="-1" style="outline: 0px;"> 
   <div class="container"> 
    <h1>Golang标准库文档</h1> 
<form action="s.php" method="get">
<input type="text" name="s"><input type="submit" value="搜索">
</form>
	<div>测试1
	测试2
	测试12
	</div>
	<div>测试3</div>
	<div>测试4</div>
	<div>测试5</div>

`

	//解释正则表达式, + 匹配前一个字符的1次或多次
	// reg := regexp.MustCompile(`<div>(.*)</div>`)
	// 匹配换行
	reg := regexp.MustCompile(`<div>(?s:(.*?))</div>`)
	if reg == nil {
		fmt.Println("MustCompile err")
		return
	}

	//提取关键信息
	// result := reg.FindAllString(buf, -1)
	result := reg.FindAllStringSubmatch(buf, -1)

	// fmt.Println("result = ", result)

	//过滤<></>
	for _, text := range result {
		fmt.Println("text[0] = ", text[0]) //带 <></>
		fmt.Println("text[1] = ", text[1]) //不带 <></>
	}

}
