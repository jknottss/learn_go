package main

import "fmt"

func main() {
    // значение по умолчанию
    var num0 int

    // значение при инициализации
    var num1 int = 1

    //пропуск типа
    var num2 = 20
    fmt.Println(num0, num1, num2)

    //короткое объявление переменных
    num := 30
    //только для новых, переприсвоить значение так нельзя

    num += 1
    fmt.Println("+=", num)

    //++num нет
    num++
    fmt.Println("++", num)

    //camelCase - принятый стиль
    userIndex := 10

    //under_score - не принятый
    user_index := 10

    fmt.Println(userIndex, user_index)

    //объявление нескольких переменных 
    var weight, height int = 10, 20

    //присваивание в существующие
    weight, height = 11, 21

    //краткое присваивание
    //хотя-бы одна переменная должна быть новой
    weight, age := 12, 22
    fmt.Println(weight, height, age)





}
