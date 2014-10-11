package main

import (
	"fmt"
	"time"
)

var ch = make(chan string)

func main() {
	go Waiter()
	Hello()
	<-ch
}

func Waiter() {
	time.Sleep(5 * time.Second)
	fmt.Println("Wait")
	ch <- "done"
}

func Hello() {
	fmt.Println("Hello")
}
