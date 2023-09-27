package main

func main() {
	//x = 1                       // 命名变量的赋值
	//*p = true                   // 通过指针间接赋值
	//person.name = "bob"         // 结构体字段赋值
	//count[x] = count[x] * scale // 数组、slice或map的元素赋值

	//v, ok = m[key]             // map
	//v, ok = x.(T)              // 类型断言
	//v, ok = <-ch               // 通道接收

	//v = m[key]                // map查找，失败时返回零值
	//v = x.(T)                 // type断言，失败时panic异常
	//v = <-ch                  // 管道接收，失败时返回零值（阻塞不算是失败）
	//
	//_, ok = m[key]            // map返回2个值
	//_, ok = mm[""], false     // map返回1个值
	//_ = mm[""]                // map返回1个值

	// 数组赋值
	//medals := []string{"gold", "silver", "bronze"}
}

// 最大公约数
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// 斐波那契
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
