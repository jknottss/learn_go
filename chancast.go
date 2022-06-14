package main

import "fmt"

//Принимаем канал только для чтения, при попытке записи будет паника
func greet(roc <-chan string) {

	for val := range roc {
		fmt.Println("Hello " + val + "!")
	}
}

func main() {
	fmt.Println("main started")

	//создаем двунаправленный канал
	c := make(chan string)

	go greet(c)

	c <- "Кирилл"

	c <- "lol"

	fmt.Println("main stopped")
}
