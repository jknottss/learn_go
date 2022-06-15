package main

import "fmt"

func squares(c chan int) {
	for i := 0; i < 9; i++ {
		c <- i * i
	}
	// закроем канал когда счетчик будетт  == 9
	close(c)
}

func main() {
	fmt.Println("prog start")

	c := make(chan int)

	go squares(c) // запускаем горутину

	// в цикле будет периодически блокироваться и разблокироваться горутина пока канал не будет закрыт
	// c range не нужно проверять закрыт ли канал
	for val := range c {
		fmt.Println(val)
	}

	fmt.Println("prog stop")
}
