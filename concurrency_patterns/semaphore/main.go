// Вы разрабатываете систему загрузки сфайлов с удаленного сервера.
// Каждый файл загружается через отдельное соединение.
// Однако, сервер ограничивает количество одновременно активных соединений,
// и ваше приложение не должно превышать этот лимит.
// В программе нельзя просто запустить фиксированное количество горутин.
// Нужно обеспечить динамическо изменяемое органичение
// на количество одновлременно работающих соединений.

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func downloadFile(filename string) {
	fmt.Printf("Downloading %s\n", filename)
	time.Sleep(1 * time.Second)
	fmt.Printf("Finished downloading %s\n", filename)
}

const maxConnections = 3

func main() {
	semaphore := make(chan struct{}, maxConnections)

	wg := sync.WaitGroup{}
	files := []string{"file1.txt", "file2.txt", "file3.txt", "file4.txt", "file5.txt", "file6.txt", "file7.txt", "file8.txt", "file9.txt", "file10.txt"}
	for _, file := range files {
		wg.Add(1)
		semaphore <- struct{}{}
		go func() {
			defer func() {
				<-semaphore
				wg.Done()
			}()

			downloadFile(file)
		}()
	}

	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
	close(semaphore)
}
