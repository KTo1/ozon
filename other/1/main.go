// Вопросы:
//
//  1. Нужно представить что ты компилятор О_о и по шагам рассказать что тут происходит
//     пока не касаясь метода GetFile
//
//  2. Теперь мы смотрим GetFile и нужно прикинуть за какое время +- 20% этот код выполнится
//     работать это все будет примерно 5 секунд, селект блокирующий он будет ждать чтения из канала
//     контекст никогда не отменится потому, что он ТУДУ, значит будет ждать одну секунду,
//     файлов 5, значит примерно 5 секунд
//
//  3. Как сделать быстрее?
//     я вижу тут два варика,
//
//  1. через вэйт группу
//     если горутны решают разные задачи
//
//  2. через эррор группу
//     если горутины решают одну задачу и если одна упала то надо валить все
//
//     используем второй варик
//     не забываем что мапа потоко не безопасна

package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	m, err := GetFiles2(context.TODO(), "1", "2", "3", "4", "5")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(m)
	fmt.Println(time.Since(start))
}

// Пример функции которую нужно оптимизировать
func GetFiles(ctx context.Context, names ...string) (result map[string][]byte, err error) {
	if len(names) == 0 {
		return nil, nil
	}

	eg, ctx := errgroup.WithContext(ctx)
	result = make(map[string][]byte)
	mu := &sync.Mutex{}
	for _, name := range names {
		eg.Go(func() error {
			res, err := GetFile(ctx, name)
			if err != nil {
				return err
			}

			mu.Lock()
			result[name] = res
			mu.Unlock()
			return nil
		})
	}

	if err = eg.Wait(); err != nil {
		return nil, fmt.Errorf("some error : %w", err)
	}

	return result, nil
}

// Пример функции которую нужно оптимизировать
// Решение через вэйтгруппу с отменой конекста через 	once := sync.Once{}
func GetFiles2(ctx context.Context, names ...string) (result map[string][]byte, err error) {
	if len(names) == 0 {
		return nil, nil
	}

	var gerror error

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	result = make(map[string][]byte)
	mu := &sync.Mutex{}
	wg := sync.WaitGroup{}
	once := sync.Once{}

	wg.Add(len(names))
	for _, name := range names {
		go func() {
			defer wg.Done()

			res, err := GetFile(ctx, name)
			if err != nil {
				once.Do(func() {
					gerror = err
					cancel()
				})

				return
			}

			mu.Lock()
			result[name] = res
			mu.Unlock()

		}()
	}

	wg.Wait()

	if gerror != nil {
		return nil, fmt.Errorf("some error : %w", gerror)
	}

	return result, nil
}

// Пример функции которую нужно оптимизировать
// Решение через вэйтгруппу с отменой конекста через канал
func GetFiles3(ctx context.Context, names ...string) (result map[string][]byte, err error) {
	if len(names) == 0 {
		return nil, nil
	}

	errCh := make(chan error, 1)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	result = make(map[string][]byte)
	mu := &sync.Mutex{}
	wg := sync.WaitGroup{}

	wg.Add(len(names))
	for _, name := range names {
		go func() {
			defer wg.Done()

			res, err := GetFile(ctx, name)
			if err != nil {
				select {
				case errCh <- err:
					cancel()
				default:
				}

				return
			}

			mu.Lock()
			result[name] = res
			mu.Unlock()

		}()
	}

	wg.Wait()

	select {
	case err = <-errCh:
		return nil, fmt.Errorf("some error : %w", err)
	default:
		return result, nil
	}
}

// Пример функции которая относительно не долго выполняется при единичном вызове
// Но достаточно долгоесли вызывать последовательно
// предроложим, что оптимизировать в ней нечено
func GetFile(ctx context.Context, name string) ([]byte, error) {
	if name == "" {
		return nil, fmt.Errorf("file name is empty, %q", name)
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-ticker.C:
	}

	if name == "invalid" {
		return nil, fmt.Errorf("file name is invalid, %q", name)
	}

	b := make([]byte, 10)
	n, err := rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("getting file %q: %w", name, err)
	}

	return b[:n], nil
}
