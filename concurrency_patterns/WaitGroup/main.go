package main

import (
	"context"
	"errors"
)

// Есть интерфейс waiter, требуется реализовать структуру, имплементирующую этот интерфейс, которая:
// - Параллельно запускает переданные в run функции с указанным контекстом.
// - Количество параллельных запусков определяется параметром maxParallel при создании waiter через newGroup.
// - Возвращает ошибку из wait, если хотя бы одна функция из run вернула её.
// - Возвращает комбинацию ошибок от вызовов run, если несколько задач завершились с ошибками (можно использовать errors.Join).

type waiter interface {
	Wait() error
	Run(ctx context.Context, f func(ctx context.Context) error)
}

type waitGroup struct {
	// напишите ваш код здесь
}

// Метод waitGroup.Wait() должен ожидать завершения всех задач и возвращать ошибку
func (g *waitGroup) Wait() error {
	// напишите ваш код здесь
}

// Метод waitGroup.Run() должен запускать функцию f с контекстом ctx
func (g *waitGroup) Run(ctx context.Context, fn func(ctx context.Context) error) {
	// напишите ваш код здесь
}

// Функция newGroupWait создает новый объект waiter с заданным количеством параллельных запусков
func newGroupWait(maxParallel int) waiter {
	// напишите ваш код здесь
}

func main() {
	g := newGroupWait(2)
	ctx := context.Background()

	expErr1 := errors.New("got error 1")
	expErr2 := errors.New("got error 2")

	g.Run(ctx, func(ctx context.Context) error {
		// Ваш код здесь
		return nil
	})

	g.Run(ctx, func(ctx context.Context) error {
		return expErr2
	})

	g.Run(ctx, func(ctx context.Context) error {
		// Ваш код здесь
		return nil
	})

	err := g.Wait()

	// Проверка на наличие ошибок expErr1 и expErr2
	if !errors.Is(err, expErr1) || !errors.Is(err, expErr2) {
		panic("wrong code")
	}

}
