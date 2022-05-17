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

func Buy(p Payer) {
	switch p.(type) {
	case *Wallet:
		fmt.Println("Оплата наличными")
	case *Card:
		plasticCard, ok := p.(*Card)
		if !ok {
			fmt.Println("Не удалось преобразовать к типу *Card")
		}
		fmt.Println("Вставляйте карту", plasticCard.CardHolder)
	default:
		fmt.Println("something new")
	}
}
func main() {
	myWallet := &Wallet{Cash: 100}
	Buy(myWallet)

	var myMoney Payer
	myMoney = &Card{Balance: 100, CardHolder: "kirill"}
	Buy(myMoney)

	myMoney = &ApplePay{Money: 9}
	Buy(myMoney)

}
