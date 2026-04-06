// Вопросы:
//  1. Нужно представить что ты компилятор О_о и по шагам рассказать что тут происходит
//     пока не касаясь метода GetFile
//  2. Теперь мы смотрим GetFile и нужно прикинуть за какое время +- 20% этот код выполнится
//		работать это все будет примерно 5 секунд, селект блокирующий он будет ждать чтения из канала
//		контекст никогда не отменится потому, что он ТУДУ, значит будет ждать одну секунду,
//		файлов 5, значит примерно 5 секунд
//	3. Как сделать быстрее?
//		я вижу тут два варика,
//		1. через вэйт группу
//			если горутны решают разные задачи
//		2. через эррор группу
//			если горутины решают одну задачу и если одна упала то надо валить все

package main

//import (
//	"context"
//	"fmt"
//	"log"
//	"math/rand"
//	"time"
//)
//
//func main() {
//	start := time.Now()
//	m, err := GetFiles(context.TODO(), "1", "2", "3", "4", "5")
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	fmt.Println(m)
//	fmt.Println(time.Since(start))
//}
//
//// Пример функции которую нужно оптимизировать
//func GetFiles(ctx context.Context, names ...string) (result map[string][]byte, err error) {
//	if len(names) == 0 {
//		return nil, nil
//	}
//
//	result = make(map[string][]byte)
//	for _, name := range names {
//		result[name], err = GetFile(ctx, name)
//		if err != nil {
//			return nil, err
//		}
//	}
//
//	return result, nil
//}
//
//// Пример функции которая относительно не долго выполняется при единичном вызове
//// Но достаточно долгоесли вызывать последовательно
//// предроложим, что оптимизировать в ней нечено
//func GetFile(ctx context.Context, name string) ([]byte, error) {
//	if name == "" {
//		return nil, fmt.Errorf("file name is empty, %q", name)
//	}
//
//	ticker := time.NewTicker(time.Second)
//	defer ticker.Stop()
//
//	select {
//	case <-ctx.Done():
//		return nil, ctx.Err()
//	case <-ticker.C:
//	}
//
//	if name == "invalid" {
//		return nil, fmt.Errorf("file name is invalid, %q", name)
//	}
//
//	b := make([]byte, 10)
//	n, err := rand.Read(b)
//	if err != nil {
//		return nil, fmt.Errorf("getting file %q: %w", name, err)
//	}
//
//	return b[:n], nil
//}
