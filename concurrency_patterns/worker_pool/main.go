// Вы разрабатываете сервис, который обрабатывает изображения.
// Каждое изображение проходит дорогостоящую обработку, например, водяной знак
// Польскольку обработка каждого изображения занимает значительное время
// необходимо обрабатывать их паралленльно, что бы ускорить процесс
// однако, что бы избежать лишней нагрузки на систему, вы хотите ограничить
// количество одновлременно работающих горутин
// Можно еще пробросить контекст

package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID   int
	file string
}

func processTask(task Task) string {
	time.Sleep(time.Second)
	return fmt.Sprintf("Processing image, %d, file - %s", task.ID, task.file)
}

func worker(taskChan <-chan Task, result chan<- string) {
	for task := range taskChan {
		result <- processTask(task)
	}
}

func generator() chan Task {
	ch := make(chan Task)

	go func() {
		defer close(ch)

		for i := range 100 {
			ch <- Task{
				ID:   i,
				file: fmt.Sprintf("%d.txt", i),
			}
		}
	}()

	return ch
}

func main() {
	ch := generator()

	maxParallel := 3
	result := make(chan string)

	wg := sync.WaitGroup{}
	for range maxParallel {
		wg.Add(1)
		go func() {
			defer wg.Done()

			worker(ch, result)
		}()
	}

	go func() {
		defer close(result)
		wg.Wait()
	}()

	start := time.Now()
	for v := range result {
		fmt.Println(v)
	}
	fmt.Println(time.Since(start))
}
