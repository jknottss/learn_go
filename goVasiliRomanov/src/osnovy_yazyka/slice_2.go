package main

import "fmt"

func main() {
	buf := []int{1, 2, 3, 4, 5}
	fmt.Println(buf)

	//получение среза указывающего на ту же область памяти
	sl1 := buf[1:4] // [2 3 4]
	sl2 := buf[:2]  // [1 2]
	sl3 := buf[2:]  // [3 4 5]

	fmt.Println(sl1, sl2, sl3)

	newBuf := buf[:] // [1 2 3 4 5]

	newBuf[0] = 9
	//теперь buf = [9 2 3 4 5]

	//теперь указывает на другую область памяти
	newBuf = append(newBuf, 6)

	//buf = [9 2 3 4 5]
	//newBuf = [1 2 3 4 5]
	newBuf[0] = 1
	fmt.Println("buf", buf)
	fmt.Println("newBuf", newBuf)

	//копирование одного слайса в другой
	var emptyBuf []int //len=0, cap=0
	//неправильно - скопирует меньшее (len) из 2х
	copied := copy(emptyBuf, buf) //copied = 0
	fmt.Println(copied, emptyBuf)

	//как нужно копировать
	newBuf = make([]int, len(buf), len(buf))
	copy(newBuf, buf)
	fmt.Println(newBuf)

	//копирование в часть существующего слайса
	ints := []int{1, 2, 3, 4}
	copy(ints[1:3], []int{5, 6}) //ints = [1 5 6 4]
	fmt.Println(ints)
}
