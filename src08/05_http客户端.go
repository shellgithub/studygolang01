package main

import (
	"fmt"
	"net/http"
)

func main() {
	// resp, err := http.Get("https://www.baidu.com")
	resp, err := http.Get("http://127.0.0.1:8000")
	if err != nil {
		fmt.Println("Listen.Get err = ", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("resp.Method = ", resp.Status)
	fmt.Println("resp.URL = ", resp.StatusCode)
	fmt.Println("resp.Header = ", resp.Header)
	fmt.Println("resp.Body = ", resp.Body)

	buf := make([]byte, 4*1024)
	var tmp string

	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("Read err = ", err)
			break
		}
		tmp += string(buf[:n])
	}
	fmt.Println("tmp = ", tmp)
}
