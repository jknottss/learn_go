package main

import "fmt"

func main() {
	fmt.Println("main started")

	c := make(chan string)

	go func(c chan string) {
		fmt.Println("Hello " + <-c + "!")
	}(c)

	c <- "Кирилл"

	fmt.Println("main stopped")
}
