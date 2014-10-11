package main

import (
	"fmt"
	"strconv"
	"time"
)

func Bakery(store chan string) {
	for i := 1; i <= 10; i++ {
		breadName := "bread " + strconv.Itoa(i)
		fmt.Println(breadName + " shipped to store!")
		store <- breadName
	}
	close(store)
}

func Consumer(store chan string) {
	for {
		bread, ok := <-store
		if !ok {
			break
		}
		fmt.Println("baught " + bread)
	}
}

func main() {
	store := make(chan string)
	fmt.Println("sotre open!")
	go Bakery(store)
	go Consumer(store)

	time.Sleep(60 * time.Second)
}
