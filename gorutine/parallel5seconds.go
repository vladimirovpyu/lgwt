package main

import (
    "fmt"
    "time"
    "math/rand"
    "context"
)

func processData(val int) int {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	return val * 2
}

func main() {
	in := make(chan int)
	out := make(chan int)

	go func() {
		for i := range 100 {
			in <- i
		}
		close(in)
	}()

	now := time.Now()
	processParallel(in, out, 5)

	for val := range out {
		fmt.Println(val)
	}
	fmt.Println(time.Since(now))
}

// операция должна выполняться не более 5 секунд
func processParallel(in <-chan int, out chan<- int, numWorkers int) {
    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)

    for i:=0; i < numWorkers; i++ {
        go func() {
            for {
                select {
                    case <- ctx.Done():
                        return
                    case val, ok := <-in:
                        if !ok {
                            return
                        }

                        res := processData(val)

                        select {
                            case <-ctx.Done():
                                return
                            case out <- res:
                        }

                }
            }
        }()
    }

	go func() {
		<- ctx.Done()
		cancel() // Освобождаем ресурсы контекста
		close(out)
	}()
}
