package main

import (
	"fmt"
)

func main() {
	const PI = 3.1415
	var r = 5.0
	var c = 2 * PI * r

	// 也可以用括号统一声明多个变量
	var (
		a = 1
		b = 2
	)

	var d int
	var (
		e, f int
	)

	// 定义后再初始化
	d = 1
	e = 2
	f = d + e

	// 重新赋值
	e = 5
	d = e + f

	fmt.Print("Hello,")
	fmt.Print("world")

	fmt.Println()

	fmt.Println("Hello,")
	fmt.Println("world")
	fmt.Println(a, b, c, d, e, f)

	fmt.Println("Hello,", "world")

	fmt.Printf("123456789\n")
	fmt.Printf("%4v|\n", 6)
	fmt.Printf("%-4v|\n", 6)
}
