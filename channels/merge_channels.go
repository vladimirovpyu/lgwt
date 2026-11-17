package main

import (
	"fmt"
	"sync"
)

func mergeChannels(cs ...<-chan int) chan int {
	out := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(len(cs))

	for _, c := range cs {
		go func(ch <-chan int) {
            defer wg.Done()
            for data := range ch {
                out <- data
            }
		}(c)
	}

	// Горутина для ожидания завершения всех воркеров и закрытия выходного канала
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func worker(c <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range c {
		out <- data
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Запись в небуферизованные каналы в отдельных горутинах
	go func() {
		ch1 <- 1
		close(ch1)
	}()

	go func() {
		ch2 <- 2
		close(ch2)
	}()

	out := mergeChannels(ch1, ch2)

	// Теперь этот цикл корректно завершится, когда канал out закроется
	for data := range out {
		fmt.Println(data)
	}
}
