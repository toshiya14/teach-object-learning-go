package main

import (
	"fmt"
)

func main(){
	arr := [4]int {1, 2, 3, 4}
	slc := arr[1:3] // -> [2, 3]
	slc2 := slc[:]
	slc2 = append(slc2, 5)
	fmt.Println("slc:", slc)
	fmt.Println("slc2:", slc2)
	fmt.Println("arr:", arr)
	slc = append(slc, 6)
	fmt.Println("slc:", slc)
	fmt.Println("slc2:", slc2)
	fmt.Println("arr:", arr)
}