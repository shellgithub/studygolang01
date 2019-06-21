package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url) //发送get请求
	if err1 != nil {
		fmt.Println("http.Get err = ", err1)
		return
	}

	defer resp.Body.Close()

	//读取网页内容
	buf := make([]byte, 4*1024)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		result += string(buf[:n]) //累加读取内容
	}
	return
}

//开始爬取每一个笑话，每一个段子 title, content, err := SpiderOneJoy(url)
func SpiderOneJoy(url string) (title, content string, err error) {
	//开始爬取页面内容
	result, err1 := HttpGet(url)
	if err1 != nil {
		fmt.Println("HttpGet err = ", err1)
		err = err1
		return
	}

	//取关键信息
	//取段子标题 <h1>段子标题</h1> 只取一个
	re1 := regexp.MustCompile(`<h1>(?s:(.*?))</h1>`)
	if re1 == nil {
		err = fmt.Errorf("%s", "regexp.MustCompile err")
		return
	}

	//取段子标题
	tmpTitle := re1.FindAllStringSubmatch(result, 1) //最后一个参数为1，只过滤第一个
	for _, data := range tmpTitle {
		title = data[1]
		// title = strings.Replace(title, "\n", "", -1)
		// title = strings.Replace(title, "\r", "", -1)
		title = strings.Replace(title, "\t", "", -1)
		break
	}

	//取段子内容 <div class="content-txt pt10"> 段子内容 <a id="prev" href=
	re2 := regexp.MustCompile(`<div class="content-txt pt10">(?s:(.*?))<a id="prev" href=`)
	if re2 == nil {
		err = fmt.Errorf("%s", "regexp.MustCompile err")
		return
	}

	//取段子内容
	tmpContent := re2.FindAllStringSubmatch(result, -1)
	for _, data := range tmpContent {
		content = data[1]
		content = strings.Replace(content, "\t", "", -1)
		content = strings.Replace(content, " ", "", -1)
		content = strings.Replace(content, "<br />", "", -1)
		content = strings.Replace(content, "&nbsp;", "", -1)
		break
	}
	return
}

//把内容写入到文件	storeJoyToFile(i, fileTitle, fileContent)
func storeJoyToFile(i int, fileTitle, fileContent []string) {
	//新建文件
	f, err := os.Create(strconv.Itoa(i) + ".txt")
	if err != nil {
		fmt.Println("os.Create err = ", err)
		return
	}

	defer f.Close()

	//写内容
	n := len(fileTitle)
	for i := 0; i < n; i++ {
		//写标题
		f.WriteString(fileTitle[i] + "\n")
		//写内容
		f.WriteString(fileContent[i] + "\n\n")
		f.WriteString("===========================================\n")
	}

}

func SpiderPage(i int, page chan int) {
	//明确爬取的url
	// url := "https://www.pengfu.com/xiaohua_1.html"
	url := "https://www.pengfu.com/xiaohua_" + strconv.Itoa(i) + ".html"
	fmt.Printf("正在爬取第%d个网页：%v\n", i, url)

	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err = ", err)
		return
	}
	// fmt.Println("r = ", result)

	//取，<h1 class="dp-b"><a href=" 一个段子的url连接 "
	//解释表达式
	re := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`)
	if re == nil {
		fmt.Println("regexp.MustCompile err = ", err)
		return
	}
	//取关键信息
	joyUrls := re.FindAllStringSubmatch(result, -1)
	// fmt.Println("joyUrls = ", joyUrls)

	fileTitle := make([]string, 0)
	fileContent := make([]string, 0)

	//取地址
	//第一个返回下标，第二个返回内容
	for _, data := range joyUrls {
		// fmt.Println("url = ", data[1])

		//开始爬取每一个笑话，每一个段子
		title, content, err := SpiderOneJoy(data[1])
		if err != nil {
			fmt.Println("SpiderOneJoy err = ", err)
			return
		}
		// fmt.Println("title = ", title)
		// fmt.Println("content = ", content)

		fileTitle = append(fileTitle, title)
		fileContent = append(fileContent, content)
	}

	// fmt.Println("fileTitle = ", fileTitle)
	// fmt.Println("fileContent = ", fileContent)

	//把内容写入到文件
	storeJoyToFile(i, fileTitle, fileContent)

	page <- i //写内容，写num

}

func DoWork(start, end int) {
	fmt.Printf("准备爬取第%d页到%d页的网址\n", start, end)

	page := make(chan int)

	for i := start; i <= end; i++ {
		//定义一个函数，爬主页面
		go SpiderPage(i, page)

	}

	for i := start; i <= end; i++ {
		//定义一个函数，爬主页面
		fmt.Printf("第%d个页面爬取完成\n", <-page)

	}

}

func main() {
	var start, end int
	fmt.Printf("请输入起始页( >= 1 ) :")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页( >= 起始页 ) :")
	fmt.Scan(&end)

	DoWork(start, end)

}
