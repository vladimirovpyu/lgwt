package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Создаем контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond)
	defer cancel()

	ch := make(chan int)

	go func() {
		// Обязательно закрываем канал при выходе из горутины
		defer close(ch)

		for i := range 1000 {
			select {
			case ch <- i:
			case <-ctx.Done():
				// Выходим из функции горутины, если контекст отменен
				return
			}
		}
	}()

	for {
		select {
		case v, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println(v)
		case <-ctx.Done():
			return
		}
	}
}

// пробрасываем контекст и нужно следить чтобы в горутинах был выход по контексту.