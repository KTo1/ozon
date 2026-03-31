// Нужно исправить код
//  1. Реализации интерфейсов работали корректно
//     объявить указателем
//     creditCard := CreditCardProcessor{limit: 100.0}
//     payPal := PayPalProcessor{balance: 200.0}
//
// у ресиверов проставить указатели
// 2. Программа завершалась успешно и без паники
// 3. Лишние вызовы методов устранить
// 4. Вы смогли бы объяснить где были ошибки и почему они возникли
// работает так потому, что если ресивер указатель то соотвествовать интерфесу будет тоже только указатель
package main

import (
	"errors"
	"fmt"
)

type PaymentProcessor interface {
	Process(amount float64) error
	Verify(amount float64) bool
}

// Реализация для кредитной карты
type CreditCardProcessor struct {
	limit float64
}

func (c *CreditCardProcessor) Process(amount float64) error {
	if amount > c.limit {
		return errors.New("amount exceeds limit credit card")
	}

	c.limit -= amount

	fmt.Printf("Processed payment of $%.2f using credic card\n", amount)
	return nil
}

func (c *CreditCardProcessor) Verify(amount float64) bool {
	return amount <= c.limit
}

// Реализация для PayPal
type PayPalProcessor struct {
	balance float64
}

func (p *PayPalProcessor) Process(amount float64) error {
	if amount > p.balance {
		return errors.New("amount exceeds limit paypal")
	}

	p.balance -= amount

	fmt.Printf("Processed payment of $%.2f using pay pal\n", amount)
	return nil
}

func (p *PayPalProcessor) Verify(amount float64) bool {
	return amount <= p.balance
}

func ExecutePayment(processor PaymentProcessor, amount float64) {
	if processor.Verify(amount) {
		err := processor.Process(amount)
		if err != nil {
			fmt.Printf("Error processing payment: %v\n", err)
		}
	} else {
		fmt.Printf("Verivication failed for amount:", amount)
	}
}

func main() {
	creditCard := &CreditCardProcessor{limit: 100.0}
	payPal := &PayPalProcessor{balance: 200.0}

	ExecutePayment(creditCard, 50)
	ExecutePayment(creditCard, 50)
	ExecutePayment(creditCard, 50)
	ExecutePayment(payPal, 50)
	ExecutePayment(payPal, 50)
}
