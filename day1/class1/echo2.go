package main

import (
	"fmt"
	"os"
)

func main() {
	// 使用 range 遍历
	s, sep := "", ""
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
