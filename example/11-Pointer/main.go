package main

import "fmt"

func add1(a int) {
	a = a + 1
	fmt.Printf("a(address): %v, a: %v\n", &a, a)
}

func add2(ap *int) {
	*ap = *ap + 1
	fmt.Printf("ap(address): %v, ap: %v\n", ap, *ap)
}

func main() {
	var num int = 1
	fmt.Printf("num(address): %v, num: %v\n", &num, num)
	add1(num)
	fmt.Printf("num: %v\n", num)
	add2(&num)
	fmt.Printf("num: %v\n", num)
}
