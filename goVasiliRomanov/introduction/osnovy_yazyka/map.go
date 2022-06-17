package main

import "fmt"

func main() {
	//инициализация при создании
	var user map[string]string = map[string]string{
		"name":     "Vasiliy",
		"lastName": "Romanov",
	}

	//с нужной емкостью
	profile := make(map[string]string, 10)

	//узнаем количество элементов
	mapLength := len(user)

	fmt.Printf("%d %+v\n", mapLength, profile)

	//если не найдет в мапе значение, вернет по умолчанию для типа
	mName := user["meddleName"]
	fmt.Println("mName:", mName)

	//проверка на существование ключа
	mName, mmNameExist := user["middleName"]
	fmt.Println("mName:", mName, "mNameExist:", mmNameExist)

	//можно не использовать значение, а только проверять есть ли ключ
	_, mmmNameExist2 := user["middleName"]
	fmt.Println("mNameExist2:", mmmNameExist2)

	//добавление элемента в мапу
	user["middleName"] = "Olegovich"

	//удаление ключа
	delete(user, "lastName")
	fmt.Printf("%#v\n", user)
}
