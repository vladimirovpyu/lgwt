package tests

import (
	"fmt"
	"testing"
	"time"
)

func start() {
	ch := make(chan int, 100)

	for i := range 100 {
		ch <- i
	}

	for range 3 {
		go worker(ch)
	}

	close(ch)

	time.Sleep(5 * time.Second)
}

func worker(ch chan int) {

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			fmt.Printf("worker tick\n")
		case c := <-ch:
			fmt.Printf("worker process channel %d\n", c)
			if c == 0 {
                fmt.Printf("worker stop\n")
                return
            }
        }
	}
}

func TestStart(t *testing.T) {
	start()
}
