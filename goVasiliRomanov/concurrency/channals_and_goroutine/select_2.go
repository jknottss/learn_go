package main

import "fmt"

func main() {
	ch1 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2
	ch2 := make(chan int, 2)
	ch2 <- 3

loop:
	for {
		select {
		case v1 := <-ch1:
			fmt.Println("chan1 val", v1)
		case v2 := <-ch2:
			fmt.Println("chan2 val", v2)
		default:
			break loop
		}
	}
	fmt.Println("channals is empty")
}
