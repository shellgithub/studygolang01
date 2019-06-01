package main // 必须

//忽略此包
// import _ "fmt"

// func main() {

// }

// 给包名起别名
// import io "fmt"

// func main() {
// 	io.Println("this is a test")

// }

// .操作
// import . "fmt"

// func main() {
// 	Println("this is a test")
// }

/*
// 方式1
// import "fmt"  //导入包，必须使用，否则编译不过
// import "os"

//方式2,常用
import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("this is a test")
	fmt.Println("os.Args =", os.Args)
}

*/
