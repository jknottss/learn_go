package main

import "fmt"

type Wallet struct {
	Cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("Нехватает денег в кошельке\n")
	}
	w.Cash -= amount
	return nil
}

type Card struct {
	Balance    int
	ValidUntil string
	CardHolder string
	CVV        string
	Number     string
}

func (c *Card) Pay(amount int) error {
	if c.Balance < amount {
		return fmt.Errorf("Нехватает денег на карте\n")
	}
	c.Balance -= amount
	return nil
}

type ApplePay struct {
	Money   int
	AppleId string
}

func (a *ApplePay) Pay(amount int) error {
	if a.Money < amount {
		return fmt.Errorf("Нехватает денег на аккаунте\n")
	}
	return nil
}

type Payer interface {
	Pay(int) error
}
//Теперь сюда передается пустой интерфейс и в рантайме(а не при компиляции) будет проверятся
//соответствует ли переданный объект нужному интерфейсу
func Buy(in interface{}) {
	var p Payer
    var ok bool
	if p, ok = in.(Payer); !ok {
		fmt.Printf("%T: не является платежным средством \n", in)
		return
	}
    err := p.Pay(10)
    if err != nil {
        fmt.Printf("Ошибка при оплате %T: %v\n")
    }
	fmt.Printf("Покупка через %T прошла успешно!\n", p)
}

func main() {
	myWallet := &Wallet{Cash: 100}
    Buy(myWallet)
    Buy([]int {1, 2, 3})
    Buy(3.14)
}
