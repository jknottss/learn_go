package main

import "fmt"

func fib(length int) <-chan int { //возвращаем канал только для чтения хотя создаем двунаправленный
	//создаем буферезированный канал
	c := make(chan int, length)

	go func() {
		for i, j := 0, 1; i < length; i, j = i+j, i {
			c <- i
		}
		close(c)
	}()

	return c
}

func main() {

	//читаем 10 чисел из канала возвращенного функцией fib()
	for fn := range fib(10) {
		fmt.Println("Current fibonacci number is", fn)
	}
}
