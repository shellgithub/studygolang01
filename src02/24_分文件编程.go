/*
1、分文件编程（多个源文件），必须放在src目录
2、设置GOPATH 环境变量
3、同一个目录，包名必须一样
src
  main.go
    package main
    test.go
    
4、 go env 查看 go 的环境路径
5、 同一个目录，调用别的文件的函数，直接调用即可，无需包名引用