package main

import "fmt"

func getSomeVars() string {
	fmt.Println("getSomeVars doing some work")
	return "result getSomeVars"
}

func main() {
	defer fmt.Println("After work")
	// defer fmt.Println(getSomeVars()) -- распечатает "getSomeVars doing some work" при инициализации (1-м)
	defer func() {
		fmt.Println(getSomeVars()) //вызовет функцию и распечатает как ожидается перед "After work"
	}()
	fmt.Println("Doing some work")
}
