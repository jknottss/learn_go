package main

import "fmt"

func greet(c chan string) {
	fmt.Println("Hello " + <-c + "!")
}

func main() {
	fmt.Println("start prog")
	c := make(chan string)

	go greet(c)

	c <- "Kirill"
	fmt.Println("prog stop")
}
