package main

import (
        "fmt"
        "strconv"
    )

type Wallet struct {
    Cash int
}

//Переопределение метода для объекта
func (w *Wallet) String() string {
    return "Мой кастомный текст " + strconv.Itoa(w.Cash) + "столько тут лежит " 
}

func main() {
    myWallet := &Wallet{Cash: 666}
    fmt.Printf("Raw payment: %#v\n", myWallet)
    fmt.Printf("Способ оплаты: %s\n", myWallet)
}
