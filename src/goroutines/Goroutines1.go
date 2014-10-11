package main

import (
	"fmt"
	"time"
)

func Hello() {
	fmt.Println("Hello")
}

func Waiter() {
	time.Sleep(5 * time.Second)
	fmt.Println("Wait")
}

func main() {
	Waiter()
	Hello()
}
