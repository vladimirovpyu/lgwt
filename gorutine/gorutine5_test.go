package tests

import (
	"fmt"
	"sync"
	"testing"
)

func start() {
	ch := make(chan int, 10)
	var wg sync.WaitGroup

    wg.Add(3)

	for i := range 10 {
		ch <- i
	}

    for i := 1; i <= 3; i++ {
		go worker(i, ch, &wg)
	}

    close(ch)

    wg.Wait()
}

func worker(id int, ch chan int, wg *sync.WaitGroup) {
    defer wg.Done()

    for c := range ch {
        fmt.Printf("%d worker process channel %d\n", id, c)
    }
}

func TestStart(t *testing.T) {
	start()
}
