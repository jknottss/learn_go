package main

import "fmt"

func main() {
	//цикл без условия как while(true) или for(;;;)
	for {
		fmt.Println("loop iteration")
		break
	}

	//цикл как while
	isRun := true
	for isRun {
		fmt.Println("loop iterations with condition")
		isRun = false
	}

	// с условием и с блоком инициализации
	for i := 0; i < 2; i++ {
		fmt.Println("loop iteration", i)
		if i == 1 {
			continue
		}
	}

	//операции по slice
	sl := []int{1, 2, 3}
	idx := 0

	for idx < len(sl) {
		fmt.Println("while-style loop, idx", idx, "value", sl[idx])
		idx++
	}

	for i := 0; i < len(sl); i++ {
		fmt.Println("c-style loop", i, sl[i])
	}
	for idx := range sl {
		fmt.Println("range slice by index", idx)
	}
	for idx, val := range sl {
		fmt.Println("range slice by index-value", idx, val)
	}
	//операции по map
	profile := map[int]string{1: "Vasily", 2: "Romanov"}

	for key := range profile {
		fmt.Println("range map by key", key)
	}
	for key, val := range profile {
		fmt.Println("range map by key-value", key, val)
	}
	for _, val := range profile {
		fmt.Println("range map by val", val)
	}

	//итерация по строке
	str := "Привет, Мир!"
	for pos, char := range str {
		fmt.Printf("%#U at pos %d\n", char, pos)
	}
}
