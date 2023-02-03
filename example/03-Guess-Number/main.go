package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var number, answer, min, max int
	min = 0
	max = 99
	rand.Seed(time.Now().UnixMilli())
	number = rand.Intn(100)
	fmt.Println("Game start")

	for {
		fmt.Printf("Input a number(%v ~ %v)", min, max)
		fmt.Scan(&answer)
		if answer < min || answer > max {
			fmt.Println("F*ck you")
			continue
		}
		if number < answer {
			fmt.Println("Large")
			max = answer - 1
		} else if number > answer {
			fmt.Println("Small")
			min = answer + 1
		} else {
			fmt.Println("Bingo")
			break
		}
	}
}
