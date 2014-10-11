package main

import "fmt"

func sum(a []int, h chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	h <- sum // send sum to c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}
	ch := make(chan int)

	go sum(a[:len(a)/2], ch)
	go sum(a[len(a)/2:], ch)
	x, y := <-ch, <-ch // receive from c

	fmt.Println(x, y, x+y)
}
