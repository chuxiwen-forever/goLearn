package main

import (
	"flag"
	"fmt"
	"strings"
)

// 命令行输入n，为bool，如果不输入为false,
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", "", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}

	//$ go build gopl.io/ch2/echo4
	//$ ./echo4 a bc def
	//a bc def
	//$ ./echo4 -s / a bc def
	//a/bc/def
	//$ ./echo4 -n a bc def
	//a bc def$
	//$ ./echo4 -help
	//Usage of ./echo4:
	//-n    omit trailing newline
	//-s string
	//separator (default " ")
}
