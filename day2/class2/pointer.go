package main

func main() {
	//x := 1
	//p := &x         // p是x的指针，指向x的内存地址
	//fmt.Println(*p) // "1"
	//*p = 2          // 使用p修改x的值
	//fmt.Println(x)  // "2"

	//var x, y int
	//fmt.Println(&x == &x, &x == &y, &x == nil) // "true false false"

	//var p = f()
	//fmt.Println(p == p)     // ture 指向地址
	//fmt.Println(f() == f()) // false 新的变量地址不一样

	//v := 1
	//incr(&v)              // side effect: v is now 2
	//fmt.Println(incr(&v)) // "3" (and v is 3)
}

func incr(p *int) int {
	*p++ // 非常重要：只是增加p指向的变量的值，并不改变p指针！！！
	return *p
}

func f() *int {
	v := 1
	return &v
}
