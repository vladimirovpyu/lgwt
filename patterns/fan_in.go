package main

import (
	"fmt"
	"sync"
	"time"
)

func case3(channels ...chan int) chan int {
	out := make(chan int)
	done := make(chan struct{}) // Сигнальный канал для остановки всех воркеров
	var closeOnce sync.Once     // Гарантирует, что done закроется ровно один раз
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c chan int) {
			defer wg.Done()
			for {
				select {
				case <-done:
					// Если другой канал уже закрылся — выходим без утечки горутины
					return
				case val, isClosed := <-c:
					if !isClosed {
						// Входной канал c закрылся — сигнализируем остальным остановиться
                        closeOnce.Do(func() {
                            close(done)
                        })
						return
					}
					// Пытаемся отправить значение в out, но не зависаем, если уже сработал done
					select {
					case out <- val:
					case <-done:
						return
					}
				}
			}
		}(ch)
	}

	// Отдельная горутина ждёт завершения всех воркеров и закрывает результирующий канал
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1
		time.Sleep(50 * time.Millisecond)
		ch1 <- 2
		close(ch1) // Закрываем первый канал — вся цепочка должна остановиться
	}()

	go func() {
		for i := 10; ; i++ {
			select {
			case ch2 <- i:
				time.Sleep(30 * time.Millisecond)
			case <-time.After(200 * time.Millisecond):
				return // Защита от зависания отправителя в примере
			}
		}
	}()

	merged := case3(ch1, ch2)
	for val := range merged {
		fmt.Println("Получено:", val)
	}
	fmt.Println("Канал merged закрыт, работа завершена")
}
