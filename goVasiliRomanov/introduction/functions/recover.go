package main

import "fmt"

func deferTest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("catch panic:", err)
		}
	}()
	fmt.Println("some work")
	panic("something bad happend")
	return
}

func main() {
	deferTest()
	return
}
