package main

import (
	"fmt"
	"sync"
)

var i int

func worker(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock() // заблокировли ресурсы для остальных потоков (горутин)
	i += 1
	m.Unlock() // разблокировали
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(&wg, &m)
	}

	wg.Wait()

	fmt.Println("Value of i after 1k iterations is", i)

}
