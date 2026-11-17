package main

import (
	"fmt"
	"sync"
)

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	// Функция-воркер для пересылки данных из одного входного канала в общий
	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan int) {
            defer wg.Done()
            for n := range c {
                out <- n
            }
        }(c)
	}

	// Запускаем отдельную горутину, которая закроет общий канал,
	// когда все входные каналы будут вычитаны и закрыты
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	// Пример использования
	asChan := func(vs ...int) <-chan int {
		c := make(chan int)
		go func() {
			for _, v := range vs {
				c <- v
			}
			close(c)
		}()
		return c
	}

	c1 := asChan(1, 2, 3)
	c2 := asChan(4, 5, 6)
	c3 := asChan(7, 8, 9)

	for v := range merge(c1, c2, c3) {
		fmt.Println(v)
	}
}
