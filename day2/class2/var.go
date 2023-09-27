package main

import (
	"fmt"
	"os"
)

func main() {
	var s1 string
	fmt.Println(s1) // ""

	var i, j, k int
	fmt.Println(i, j, k) // int int int
	var b, f, s2 = true, 2.3, "four"
	fmt.Println(b, f, s2) // bool float64 string

	var f1, err = os.Open("") // 函数返回两个返回值
	if err == nil {
		fmt.Println(f1)
	}

	// 简短变量声明
	// 0 无限循环，其他指定次数
	//anim := gif.GIF{LoopCount: 1}
	//freq := rand.Float64() * 3.0
	//t := 0
	//i := 100                  // an int
	//var boiling float64 = 100 // a float64
	//var names []string
	//var err error
	//var p Point

	//i, j := 0, 1
	//i, j := j, i 交换i和j的值
}
