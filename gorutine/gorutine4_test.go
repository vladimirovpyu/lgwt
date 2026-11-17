package tests

import (
	"fmt"
	"testing"
)

func start() {
	ch := make(chan int, 10)
	done := make(chan bool)

	for i := range 10 {
		ch <- i
	}

    for i := 1; i <= 3; i++ {
		go worker(i, ch, done)
	}

    close(ch)

    for i := 0; i < 3; i++ {
        <-done // Блокировка: ждем тру от очередного воркера
    }
}

func worker(id int, ch chan int, done chan bool) {
    for c := range ch {
        fmt.Printf("%d worker process channel %d\n", id, c)
    }

    done <- true
}

func TestStart(t *testing.T) {
	start()
}
