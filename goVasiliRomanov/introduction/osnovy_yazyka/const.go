package main

import "fmt"

const pi = 3.141
const (
	hello = "Привет"
	e     = 2.718
)
const (
	zero  = iota
	_     //пустая переменная
	three // = 3
)
const (
	_         = iota             //пропускаем первое значение
	KB uint64 = 1 << (10 * iota) // 1042
	MB                           // 1048576
)

const (
	//нетипизированная константа
	year = 2017
	//типизированная
	yearTyped int = 2017
)

func main() {
	var month int32 = 13
	fmt.Println(month + year) // а так можно потому что нетипизированная

	//mismatched types int32 and int
	//fmt.Println(month + yearTyped)
}
