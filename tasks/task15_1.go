package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Её тело нельзя изменять
func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000)
	time.Sleep(time.Duration(rnd) * time.Millisecond)
	return rnd
}

// Обертка с таймаутом в 1 секунду
func predictableFunc() (int64, error) {
	start := time.Now()
	defer func() {
		log.Printf("Время выполнения: %s\n", time.Since(start))
	}()

	// Буферизованный канал на 1 элемент, чтобы горутина не заблокировалась навечно,
	// если мы уйдем по таймауту раньше, чем она завершится.
	ch := make(chan int64, 1)

	go func() {
		ch <- unpredictableFunc()
	}()

	select {
	case res := <-ch:
		return res, nil
	case <-time.After(1 * time.Second):
		return 0, errors.New("timeout: функция выполнялась слишком долго")
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
