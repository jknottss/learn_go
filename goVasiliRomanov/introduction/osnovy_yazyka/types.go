package main

type UserID int

func main() {
	idx := 1
	var uid UserID = 42

	//даже если базовый тип один, разные типы не совместимы
	//myID := idx - так нельзя

	myID := UserID(idx)

	println(uid, myID)
}
