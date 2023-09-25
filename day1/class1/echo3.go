package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 使用 strings 内置函数
	fmt.Println(strings.Join(os.Args[1:], " "))
}
