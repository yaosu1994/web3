package main

import "fmt"

type PayMethod interface {
	Account
	Pay(int) bool
}

type Account interface {
	GetBalance() int
}

type CreditCard struct {
	balance int
	limit   int
}

func (c *CreditCard) Pay(amount int) bool {
	if c.balance+amount <= c.limit {
		c.balance += amount
		fmt.Printf("信用卡成功支付：%d\n", amount)
		return true
	} else {
		fmt.Printf("信用卡余额不足，支付失败\n")
		return false
	}
}

func (c *CreditCard) GetBalance() int {
	return c.balance
}

func (c *CreditCard) GetLimit() int {
	return c.limit
}

type BankCard struct {
	balance int
}

func (b *BankCard) Pay(amount int) bool {
	if b.balance >= amount {
		b.balance -= amount
		fmt.Printf("银行卡成功支付：%d\n", amount)
		return true
	} else {
		fmt.Printf("银行卡余额不足，支付失败\n")
		return false
	}
}

func (b *BankCard) GetBalance() int {
	return b.balance
}

func purchase(payMethod PayMethod, amount int) {
	if payMethod.Pay(amount) {
		fmt.Printf("支付成功，余额：%d\n", payMethod.GetBalance())
	} else {
		switch card := payMethod.(type) {
		case *CreditCard:
			fmt.Printf("信用卡可用额度不足，可用额度：%d，请使用其他支付方式\n", card.GetLimit()-card.GetBalance())
		case *BankCard:
			fmt.Printf("银行卡余额不足，余额：%d, 请使用其他支付方式\n", card.GetBalance())
		}
	}

}
func main() {
	creditCard := &CreditCard{balance: 0, limit: 1000}
	bankCard := &BankCard{balance: 500}

	fmt.Println("使用信用卡支付：")
	purchase(creditCard, 800)
	fmt.Println()
	purchase(creditCard, 300)
	fmt.Println()
	fmt.Println("使用银行卡支付：")
	purchase(bankCard, 300)
	fmt.Println()
	purchase(bankCard, 300)
}
