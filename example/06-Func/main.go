package main

import (
	"fmt"
)

type role int

func (me role) say() {
	if me == 1 {
		fmt.Println("我是公的。")
	} else if me == 0 {
		fmt.Println("我是母的。")
	} else {
		fmt.Println("我是沃尔玛购物袋。")
	}
}

func main(){
	var me role = 1
	me.say()
}