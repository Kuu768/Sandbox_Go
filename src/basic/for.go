package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// whileっぽいforループ（goではwhileは使えない）
	counter1 := 0
	for counter1 < 10 {
		fmt.Println(counter1)
		counter1++
	}

	// 無限ループ
	counter2 := 0
	for {
		fmt.Println("loop")
		counter2++

		if counter2 == 10 {
			return
		}
	}
}
