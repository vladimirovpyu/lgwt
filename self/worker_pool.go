/**
 Worker Pool с ограничением параллелизма и Graceful Shutdown
	• Суть: Написать пул воркеров, который читает задачи из канала, обрабатывает их в N горутинах и
	 корректно завершается по отмене контекста или закрытию входного канала, дожидаясь завершения текущих задач.
	• Что проверяют: sync.WaitGroup, select, закрытие каналов без паники (send on closed channel), обработку ctx.Done().
*/

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func WorkerPool(ctx context.Context, workers int, jobs <-chan int) <-chan int {
	result := make(chan int)
	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case job, ok := <-jobs:
					if !ok {
						return
					}
					result <- job * 2
					time.Sleep(100 * time.Millisecond)
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()
	workers := 3
	jobs := make(chan int)
	sum := 0

	go func() {
		defer close(jobs)
		for i := 1; i <= 10000; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("Done")
				return
			case jobs <- 1:
			}
		}
	}()

	for res := range WorkerPool(ctx, workers, jobs) {
		sum += res
		fmt.Println(res)
	}
	fmt.Println("sum:", sum)
}
