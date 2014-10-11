package main

import (
	"fmt"
)

func main() {
	fmt.Println("Let's start FizzBuzz!\n")

	for i := 1; i < 30; i++ {
		str := "--"
		if i%(3*5) == 0 {
			str = "FizzBuzz!"
		} else if i%5 == 0 {
			str = "Buzz"
		} else if i%3 == 0 {
			str = "Fizz"
		} else {
			// Do Nothing
		}
		fmt.Printf("%d\t: %s\n", i, str)
	}
}
