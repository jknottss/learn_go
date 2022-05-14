package main

import "fmt"

type Person struct {
    Id      int
    Name    string
}

//не изменит оригинальной структуры, для которой вызван
func (p Person) UpdateName(name string) {
    p.Name = name
}

//изменит оригинальную структуру
func (p *Person) SetName(name string) {
    p.Name = name
}

//будет иметь больштй приоритет и поменяет поле Имя в структуре Аккаунт
func (a *Account) SetName(name string){
    a.Name = name
}

//методы могут быть не только у структуры
type MySlice []int

//добавление значения в слайс
func (m *MySlice) Add(val int) {
    *m = append(*m, val)
}

//получение количества значений
func (m *MySlice) Count() int {
    return len(*m)
}

type Account struct {
    Id      int
    Name    string
    Person
}

func main() {
    pers := Person{1, "Kirill"}
    pers.SetName("Kirill Klimov")
    fmt.Printf("update person: %#v\n", pers)

    var acc Account = Account{
        Id: 1,
        Name: "kirill",
        Person: Person{
            Id: 2,
            Name: "Kirill Klimov",
        },
    }
    acc.SetName("klimov kirill acc")
    acc.Person.SetName("kirill_klimov_person")
    fmt.Printf("update acc: %#v\n", acc)

    s1 := MySlice([]int{1, 2, 3})
    s1.Add(666)
    fmt.Println(s1.Count(), s1)

}




