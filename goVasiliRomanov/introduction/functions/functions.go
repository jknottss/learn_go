package main

import "fmt"

//обычное объявление
func singleIn(in int) int {
	return in
}

//много параметров
func multIn(a, b int, c int) int {
	return a + b + c
}

//именованый результат
func nameReturn() (out int) {
	out = 2
	return
}

//несколько результатов
func multipleReturn(in int) (int, error) {
	if in > 2 {
		return 0, fmt.Errorf("some error")
	}
	return in, nil
}

//несколько именованных результатов
func multipleNamedReturn(ok bool) (rez int, err error) {
	rez = 1
	if ok {
		err = fmt.Errorf("some error")
		//аналогично return rez, err
		//или return 1, fmt.Errorf("some error")
		return
	}
	rez = 2
	return
}

//не фиксированное количество параметров
//в этом случае в переменную in придет слайс интов
func sum(in ...int) (result int) {
	fmt.Printf("in := %#v\n", in)
	for _, val := range in {
		result += val
	}
	return
}
func main() {

	fmt.Println(multipleNamedReturn(true))

	nums := []int{1, 2, 3, 4}
	fmt.Println(nums, sum(nums...))
	return

}
