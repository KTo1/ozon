package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
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

var _ waiter = (*waitGroup)(nil)

type waitGroup struct {
	maxParallel int
	mu          sync.Mutex
	wg          sync.WaitGroup
	sem         chan struct{}
	errs        []error
}

// Функция newGroupWait создает новый объект waiter с заданным количеством параллельных запусков
func newGroupWait(maxParallel int) waiter {
	return &waitGroup{
		maxParallel: maxParallel,
		sem:         make(chan struct{}, maxParallel),
	}
}

// Метод waitGroup.Wait() должен ожидать завершения всех задач и возвращать ошибку
func (g *waitGroup) Wait() error {
	g.wg.Wait()

	g.mu.Lock()
	defer g.mu.Unlock()

	return errors.Join(g.errs...)
}

// Метод waitGroup.Run() должен запускать функцию f с контекстом ctx
func (g *waitGroup) Run(ctx context.Context, fn func(ctx context.Context) error) {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()

		select {
		case <-ctx.Done():
			g.AddErr(ctx.Err())
			return
		case g.sem <- struct{}{}:
		}

		defer func() {
			<-g.sem
		}()

		if err := fn(ctx); err != nil {
			g.AddErr(err)
		}
	}()
}

func (g *waitGroup) AddErr(err error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.errs = append(g.errs, err)
}

func main() {
	g := newGroupWait(2)
	ctx := context.Background()

	expErr1 := errors.New("got error 1")
	expErr2 := errors.New("got error 2")

	g.Run(ctx, func(ctx context.Context) error {
		fmt.Println("work hard 1")
		time.Sleep(time.Second)

		return nil
	})

	g.Run(ctx, func(ctx context.Context) error {
		return expErr2
	})

	g.Run(ctx, func(ctx context.Context) error {
		fmt.Println("work hard 2")
		time.Sleep(2 * time.Second)

		return expErr1
	})

	err := g.Wait()

	// Проверка на наличие ошибок expErr1 и expErr2
	if !errors.Is(err, expErr1) || !errors.Is(err, expErr2) {
		panic("wrong code")
	}

}
