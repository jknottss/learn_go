package main

import "fmt"

func main() {
	//размер массива является частью его типа

	//инициализируются значениями по умолчанию
	var a1 [3]int //[0,0,0]
	fmt.Println("a1 short", a1)
	fmt.Printf("a1 short %v\n", a1)
	fmt.Printf("a1 full %#v\n", a1)

	const size = 2
	var a2 [2 * size]bool //[false false false false]
	fmt.Println("a2", a2)

	//определение размера при объявлении
	a3 := [...]int{1, 2, 3}
	fmt.Println("a3", a3)

	//проверка при компиляции или выполнении
	//invalid array index 4
	//a3[4] = 12
}
