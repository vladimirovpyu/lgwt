package main

import (
    "fmt"
)

func main() {
	ch := make(chan int)
	for i := 0; i < 3; i++ {
		go func(v int) {
			ch <- v * v
		}(i)
	}
	var sum int
	for i := 0; i < 3; i++ {
		sum += <- ch
	}
	close(ch)
	fmt.Printf("result: %d\n", sum)
}
