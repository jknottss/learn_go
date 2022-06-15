package main

import (
	"fmt"
	"sync"
	"time"
)

func service(wg *sync.WaitGroup, instance int) {
	time.Sleep(2 * time.Second)
	fmt.Println("Service called on instance", instance)
	wg.Done() //убавляем счетчик (декремент) поэтому и передаем указатель
}

func main() {
	fmt.Println("main started")
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) //прибавляем счетчик (инкрементируем)
		go service(&wg, i)
	}

	wg.Wait() //тут блокируем пока счетчик не станет равен 0
	fmt.Println("main stopped")
}
