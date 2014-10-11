package main

import (
	"fmt"
)

func add1(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

// 複数の戻り値を返す関数
func add_and_minus(x, y int) (int, int) {
	return x + y, x - y
}

func main() {
	fmt.Println(add1(100, 200))
	fmt.Println(add2(100, 200))
	fmt.Println(add_and_minus(100, 200))
}
