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
	err := p.Pay(10)
	if err != nil {
		fmt.Printf("Ошибка при оплате %T: %v\n", p, err)
		return
	}
	fmt.Printf("Покупка через %T прошла успешно!\n", p)
}

func main() {
	myWallet := &Wallet{Cash: 100}
	Buy(myWallet)

	var myMoney Payer
	myMoney = &Card{Balance: 100}
	Buy(myMoney)

	myMoney = &ApplePay{Money: 9}
	Buy(myMoney)

}
