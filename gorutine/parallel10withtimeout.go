package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// longTask имитирует долгую работу с поддержкой отмены через контекст.
func longTask(ctx context.Context, workerID int) error {
	// Случайное время выполнения от 1 до 5 секунд
	workDuration := time.Duration(1+rand.Intn(5)) * time.Second

	select {
	case <-time.After(workDuration):
		fmt.Printf("Воркер #%d: успешно завершил работу за %v\n", workerID, workDuration)
		return nil
	case <-ctx.Done():
		// Срабатывает, если время вышло (таймаут)
		fmt.Printf("Воркер #%d: прерван по таймауту (%v)\n", workerID, ctx.Err())
		return ctx.Err()
	}
}

// runWorkers запускает указанное количество горутин с общим таймаутом.
func runWorkers(numWorkers int, timeout time.Duration) {
	// Создаем контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	// Обязательно освобождаем ресурсы контекста при выходе
	defer cancel()

	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			_ = longTask(ctx, id)
		}(i)
	}

	// Ждем, пока все горутины либо закончат работу, либо выйдут по таймауту
	wg.Wait()

	// Проверяем, по какой причине завершилось выполнение
	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("Итог: выполнение прервано, так как вышел общий таймаут.")
	} else {
		fmt.Println("Итог: все 10 горутин успешно уложились в отведенное время.")
	}
}

func main() {
	now := time.Now()
	const workersCount = 10
	const maxTimeout = 3 * time.Second

	fmt.Printf("Запускаем %d горутин с таймаутом %v...\n", workersCount, maxTimeout)
	runWorkers(workersCount, maxTimeout)
	fmt.Println(time.Since(now))
}
