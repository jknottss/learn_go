package main

import (
	"fmt"
	"sync"
	"time"
)

//пример воркера возводящего в квадрат (также использование однонаправленных каналов)
func sqrWorker(wg *sync.WaitGroup, tasks <-chan int, results chan<- int, instans int) {
	for num := range tasks {
		time.Sleep(time.Millisecond) //simulating block chan
		fmt.Printf("[worker %v] Sending result by worker %v\n", instans, instans)
		results <- num * num
	}
	wg.Done()
}

func main() {
	fmt.Println("[main] main() started")
	var wg sync.WaitGroup

	tasks := make(chan int, 10)
	results := make(chan int, 10)

	//запуск 3х горутин(воркеров)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go sqrWorker(&wg, tasks, results, i)
	}

	//запись в канал tasks 5 значений, не будет блокировки т.к буфер не заполнен
	for i := 0; i < 5; i++ {
		tasks <- i * 2
	}

	fmt.Println("[main] Wrote 5 tasks")

	//закрываем канал
	close(tasks)

	wg.Wait() //ждем завершения работы всех воркеров

	//получаем результаты
	for i := 0; i < 5; i++ {
		result := <-results //не блокируется потому что буфер не пустой
		fmt.Println("[main] Result", i, ":", result)
	}
	fmt.Println("[main] main() stopped")
}
