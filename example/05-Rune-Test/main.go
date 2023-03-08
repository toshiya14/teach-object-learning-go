package main

import (
	"fmt"
)

type sb = uint32

func main() {
	str := "hello,world"
	for i, c := range str {
		fmt.Println(i, c)
	}
}
