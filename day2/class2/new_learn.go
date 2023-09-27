package main

func main() {
	//p := new(int)   // p, *int 类型, 指向匿名的 int 变量
	//fmt.Println(*p) // "0"
	//*p = 2          // 设置 int 匿名变量的值为 2
	//fmt.Println(*p) // "2"

	//p := new(int)
	//q := new(int)
	//fmt.Println(p == q) // "false"
}

// 相似
//
//	func newInt() *int {
//		return new(int)
//	}
func newInt() *int {
	var dummy int
	return &dummy
}
