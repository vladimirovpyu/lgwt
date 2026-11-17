package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	timer := time.NewTimer(1 * time.Millisecond)

	ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond)
	defer cancel()

	select {
	case v := <-ch1:
		fmt.Println("v =", v, "from ch1")
	case v := <-ch2:
		fmt.Println("v =", v, "from ch2")
	case <-time.After(1 * time.Second):
		fmt.Println("exited by after")
	case <-timer.C:
		fmt.Println("exited by timer")
	case <-ctx.Done():
		fmt.Println("exited by context")
	}
}

// из всех этих паттернов лучше всего использовать контекст


