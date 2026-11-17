package tests

import (
	"fmt"
	"testing"
	"time"
)

func start() {
	ch := make(chan int, 10)

	for i := range 10 {
		ch <- i
	}

	for range 3 {
		go worker(ch)
	}

	close(ch)

	time.Sleep(2 * time.Second)
}

func worker(ch chan int) {
    for c := range ch {
        fmt.Printf("worker process channel %d\n", c)
    }

    fmt.Println("worker stop")
}

func TestStart(t *testing.T) {
	start()
}
