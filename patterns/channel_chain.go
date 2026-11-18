package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	out := make(chan string) // Общий выходной канал для результатов
	cnt := 5

	// Начальный сигнальный канал для первой горутины (i = 0)
	prevTurn := make(chan struct{})
	close(prevTurn) // Закрываем сразу, чтобы горутина 0 могла писать без ожидания

	for i := 0; i < cnt; i++ {
		nextTurn := make(chan struct{}) // Сигнальный канал для следующей горутины

		go func(id int, waitFor, signalNext chan struct{}) {
			// 1. Полезная работа выполняется конкурентно и с разной скоростью
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			result := fmt.Sprintf("Goroutine %d", id)

			// 2. Ждём своей очереди (пока предыдущая горутина не закроет канал)
			<-waitFor

			// 3. Пишем в общий канал строго по порядку
			out <- result

			// 4. Передаём эстафету следующей горутине
			close(signalNext)
		}(i, prevTurn, nextTurn)

		// Сдвигаем указатель для следующей итерации цикла
		prevTurn = nextTurn
	}

	// Читаем из общего канала — порядок гарантирован (0, 1, 2, 3, 4)
	for i := 0; i < cnt; i++ {
		fmt.Println(<-out)
	}
}
