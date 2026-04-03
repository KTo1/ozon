// Вы разрабатываете систему для обработки фин транзакций
// каждая транзакция проходит несколько этапов обработки:
// 1. Чтение транзакци из исходных данных
// 2. Фильтрация: убрать транзакции с отрицательными суммами
// 3. Конвернтация валюты: преобразовать сумму в доллары
// 4. Сохранение результатов: записываем обработанные транзакции в итоговый список

// Используем паттерн пайплайн, когда задача состоит из этапов и этапы определены он подходит, каждая таска проходит
// по каждому из этапов, но это не значит пайплайн работает последоватеьлно, каждая таска по пайплайну идет
// последолвательно, а вот сами таски паралелятся, т.е. пришла первая таска она пошла на этап, потом пошла на второй
// в это время пришла новая таска, она пошла на перый и т.д.

// так же можно прокинуть контекст сюда

package main

import (
	"fmt"
	"math/rand"
)

type Transaction struct {
	ID            int
	Amount        int
	AmountDollars int
}

func filter(in <-chan Transaction) <-chan Transaction {
	out := make(chan Transaction)

	go func() {
		defer close(out)

		for t := range in {
			if t.Amount > 0 {
				out <- t
			}
		}
	}()

	return out
}

func convert(in <-chan Transaction) <-chan Transaction {
	out := make(chan Transaction)

	go func() {
		defer close(out)

		for t := range in {
			t.AmountDollars = t.Amount * 10
			out <- t
		}
	}()

	return out
}

func generate(count int) <-chan Transaction {
	out := make(chan Transaction)

	go func() {
		defer close(out)

		for i := range count {
			t := Transaction{
				ID:            i,
				Amount:        rand.Intn(100) - 50,
				AmountDollars: 0,
			}

			fmt.Println("Generate transaction", t)
			out <- t
		}
	}()

	return out
}

func save(in <-chan Transaction) {
	for t := range in {
		fmt.Printf("Saving transaction %d\n", t)
	}
}

func main() {
	ch := generate(10)

	save(convert(filter(ch)))
}
