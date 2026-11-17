package tests

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func start() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan int, 100)

	close(ch)

	for i := range 100 {
		ch <- i
	}

	for chNum := range 10 {
		wctx := context.WithValue(ctx, "num", chNum)
		go worker(wctx, ch)
	}

	time.Sleep(5 * time.Second)
	cancel()
}

func worker(ctx context.Context, ch chan int) {
	wNum := ctx.Value("num").(int)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%d worker done\n", wNum)
			return
		case <-ticker.C:
			fmt.Printf("%d worker tick\n", wNum)
		case c := <-ch:
			fmt.Printf("%d worker process channel %d\n", wNum, c)
		}
	}
}

func TestStart(t *testing.T) {
	start()
}
