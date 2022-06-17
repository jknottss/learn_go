package main

import "fmt"

type Person struct {
	Id      int
	Name    string
	Address string
}

type Account struct {
	Id      int
	Name    string
	Cleaner func(string) string
	Owner   Person
    Person
}

func main() {
	//полное объявление структуры
	var acc Account = Account{
		Id:   1,
		Name: "kirill",
        Person: Person{
            Address: "barnaul",
        },
	}
	fmt.Printf("%#v\n", acc)

	//короткое объявление структуры
	acc.Owner = Person{2, "kirill", "qasakstan"}

	fmt.Printf("%#v\n", acc)
    fmt.Println(acc.Address)
}
