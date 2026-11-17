package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000)
	time.Sleep(time.Duration(rnd) * time.Millisecond)
	return rnd
}

func predictableFunc() (int64, error) {
	start := time.Now()
	defer func() {
		log.Printf("Время выполнения: %s\n", time.Since(start))
	}()

	// Создаем контекст с таймаутом 1 секунда
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel() // Обязательно освобождаем ресурсы таймера

	// Небуферизованный канал
	ch := make(chan int64)

	go func() {
		res := unpredictableFunc()
		// Пытаемся записать результат. Если таймаут уже вышел, сработает ctx.Done()
		select {
		case ch <- res:
		case <-ctx.Done():
			return
		}
	}()

	select {
	case res := <-ch:
		return res, nil
	case <-ctx.Done():
		// Вернет context.DeadlineExceeded
		return 0, ctx.Err()
	}
}

func main() {
	fmt.Println("started")
	res, err := predictableFunc()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Результат:", res)
}
