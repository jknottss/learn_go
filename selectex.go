package main

import (
	"fmt"
	"time"
)

var start time.Time

// init всегда отрабатывает самой первой ее вызывать не нужно (сразу при импорте пакета)
func init() {
	start = time.Now()
}

func service1(c chan string) {
	time.Sleep(3 * time.Second)
	c <- "Hello from service 1"
}

func service2(c chan string) {
	time.Sleep(5 * time.Second)
	c <- "Hello from service 2"
}

func main() {
	fmt.Println("main started")

	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	select {
	case res := <-chan1:
		fmt.Println("response from service 1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("response from service 2", res, time.Since(start))
	}

	fmt.Println("main stopped", time.Since(start))
}
