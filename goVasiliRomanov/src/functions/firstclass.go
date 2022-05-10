package main

import "fmt"

//обычная функция
func doNothing() {
    fmt.Println("I'm a regular function")
}

func main() {
    //анонимная функция
    func(in string) {
        fmt.Println("anon func out:", in)
    }("nobody")

    //присвоение анонимной функции в переменную
    printer := func(in string) {
        fmt.Println("printer out:", in)
    }
    printer("as variable")

    //определяем тип функции
    type strFuncType func(string)
    
    //функция принимает коллбек (принимает функцию и вызывает её)
    worker := func(callbeck strFuncType) {
        callbeck("as callbeck")
    }
    worker(printer)
    //функция возвращает замыкание
    prefixer := func(prefix string) strFuncType {
        return func (in string) {
            fmt.Printf("[%s] %s\n", prefix, in)
        }
    }
    successLogger := prefixer("SUCCESS")
    successLogger("exepted behavior")
}
