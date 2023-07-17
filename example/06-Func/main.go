package main

import (
	"fmt"
	"strconv"
)

func main() {

	a := 0
	b := 0

	func() {
		a := 1
		b := 2
		fmt.Println("a + b = " + strconv.Itoa(a+b))
	}()

	fmt.Println("a + b = " + strconv.Itoa(a+b))
}
