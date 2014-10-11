package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func run() error {
	return &MyError{time.Now(), "it didn't work"}
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}
