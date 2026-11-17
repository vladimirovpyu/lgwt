package main

import (
	"fmt"
	"sync"
)

const url = `http://jsonplaceholder.typicode.com/tools/%d`
const chunkSize = 100
const dataCount = 2 << 10 // 2048

func fetch(url string) string { return "data" }
func writeToFile(data string) {}

func main() {
	sem := make(chan struct{}, chunkSize) // Семафор на 100 слотов
	var wg sync.WaitGroup
	var fileMu sync.Mutex // Мьютекс для безопасной записи в один файл

	for i := 0; i < dataCount; i++ {
		wg.Add(1)
		sem <- struct{}{} // Блокируемся, если уже выполняется 100 запросов

		go func(id int) {
			defer wg.Done()
			defer func() { <-sem }() // Освобождаем слот после завершения

			data := fetch(fmt.Sprintf(url, id))

			fileMu.Lock()
			writeToFile(data)
			fileMu.Unlock()
		}(i)
	}

	wg.Wait()
}
