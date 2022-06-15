package main

import (
        "fmt"
        "runtime"
    )

func square(c chan int) {
	fmt.Println("[square] reading")
	num := <-c
	c <- num * num
}

func cube(c chan int) {
	fmt.Println("[cube] reading")
	num := <-c
	c <- num * num * num
}

func main() {
	fmt.Println("[main] main() started")

	squareChan := make(chan int)
	cubeChan := make(chan int)

	go square(squareChan)
	go cube(cubeChan)
    
    fmt.Println("active goroutines", runtime.NumGoroutine())
	testNum := 3
    
	fmt.Println("[main] sent testNum to squareChan")
	squareChan <- testNum
	fmt.Println("[main] resuming")

	fmt.Println("[main] sent testNumber to cubeChan")
	cubeChan <- testNum
	fmt.Println("[main] resuming")

	fmt.Println("[main] reading from channals")
	squareVal, cubeVal := <-squareChan, <-cubeChan
	fmt.Printf("squareVal - [%d], cubeVal - [%d]\n", squareVal, cubeVal)

	sum := squareVal + cubeVal
	fmt.Println("[main] sun of square and cube of", testNum, " is", sum)
	fmt.Println("[main] main() stopped")

}
