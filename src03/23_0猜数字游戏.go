package main // 必须

import (
	"fmt"
	"math/rand"
	"time"
)

func CreatNum(p *int) {
	//设置种子
	rand.Seed(time.Now().UnixNano())

	var num int
	for {
		num = rand.Intn(10000) //一定是4位数
		if num >= 1000 {
			break
		}

	}

	fmt.Println("num = ", num)
	*p = num

}

func GetNum(s []int, num int) {
	s[0] = num / 1000       //取千位
	s[1] = num % 1000 / 100 // 取百位
	s[2] = num % 100 / 10   // 取十位
	s[3] = num % 10         // 取个位
}

func OnGame(s []int) {
	var num int

	for {
		fmt.Printf("请输入一个4位数：")
		fmt.Scan(&num)

		// 999 < num << 10000
		if 999 < num && num < 10000 {

			break
		}
		fmt.Println("输入的数不符合要求\n")

	}
	fmt.Println("num = ", num)

	keySlice := make([]int, 4)
	GetNum(keySlice, num)
	fmt.Println("keySlice = ", keySlice)

	n := 0
	for i := 0; i < 4; i++ {
		if keySlice[i] > randSlice[i] {
			fmt.Printf("第%d位大了一点\n", i+1)
		} else if keySlice[i] < randSlice[i] {
			fmt.Printf("第%d位小了一点\n", i+1)
		} else {
			fmt.Printf("第%d位猜对了\n", i+1)
			n++
		}
	}

}

func main() {

	var randNum int

	//产生一个 4 位的随机数
	CreatNum(&randNum)
	fmt.Println("randNum: ", randNum)

	randSlice := make([]int, 4)
	// 保存这个4位数的每一位

	GetNum(randSlice, randNum)
	fmt.Println("randSlice = ", randSlice)

	/*  测试代码
	// 保存这个4位数的每一位
	n1 := 1234 / 1000   //取商取整
	n2 := (1234 % 1000) //取余数，结果234 , 234/100 取商得到2
	fmt.Println("n1 = ", n1)
	fmt.Println("n2 = ", n2)
	*/

	OnGame(randSlice) //游戏

}
