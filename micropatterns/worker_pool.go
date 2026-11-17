package main

import (
	"fmt"
	"sync"
)

const url = `http://jsonplaceholder.typicode.com/tools/%d`
const chunkSize = 100
const dataCount = 2 << 10 // 2048

// Заглушки функций по условию задачи
func fetch(url string) string { return "data" }
func writeToFile(data string) {}

func main() {
	jobs := make(chan string, chunkSize)
	results := make(chan string, chunkSize)

	var wg sync.WaitGroup

	// 1. Запускаем ровно chunkSize (100) воркеров
	for w := 0; w < chunkSize; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for reqURL := range jobs {
				data := fetch(reqURL)
				results <- data
			}
		}()
	}

	// 2. Отдельная горутина последовательно пишет результаты в файл
	// (один писатель гарантирует отсутствие data race при работе с файлом)
	writerDone := make(chan struct{})
	go func() {
		defer close(writerDone)
		for data := range results {
			writeToFile(data)
		}
	}()

	// 3. Наполняем очередь задач (2048 URL)
	for i := 0; i < dataCount; i++ {
		jobs <- fmt.Sprintf(url, i)
	}
	close(jobs) // Сигнал воркерам, что задачи закончились

	// 4. Ждём завершения всех воркеров, закрываем results и ждём дозаписи файла
	wg.Wait()
	close(results)
	<-writerDone
}
