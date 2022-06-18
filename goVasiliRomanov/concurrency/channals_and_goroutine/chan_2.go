package main

import "fmt"

func main() {
	in := make(chan int)

	go func(out chan<- int) {
		for i := 0; i <= 10; i++ {
			fmt.Println("before", i)
			out <- i
			fmt.Println("after", i)
		}
		close(out)
		fmt.Println("generator finish")
	}(in)

	for val := range in {
		fmt.Println("\tget", val)
	}
	fmt.Scanln()
}
