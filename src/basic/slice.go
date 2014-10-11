package main

import (
	"fmt"
)

func main() {
	list := []int{1, 2, 3, 4, 5}

	fmt.Println(list)
	fmt.Println(list[0])
	fmt.Println(list[1])
	fmt.Println(list[2])
	fmt.Println(list[3])
	fmt.Println(list[4])

	fmt.Println(list[2:])
	fmt.Println(list[:2])
	fmt.Println(list[1:3])

	// 要素の追加
	list = append(list, 6, 7, 8)
	fmt.Println(list)

	// 要素を削除したい
	list = append(list[:2], list[3:]...)
	fmt.Println(list)

	list2 := []string{"hoge", "fuga", "piyo"}

	fmt.Println(list2)
	fmt.Println(list2[0])
	fmt.Println(list2[1])
	fmt.Println(list2[2])

	// 要素の追加
	list2 = append(list2, "hoge")
	fmt.Println(list2)
}
