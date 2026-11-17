package main

import (
    "fmt"
//    "sync"
//    "time"
    )

func main() {
    done := make(chan bool, 5)
	for i := 0; i < 5; i++ {
		go func(v int, done chan bool) {
			fmt.Println(v)
			done <- true
		}(i, done)
	}

    for i := 0; i < 5; i++ {
        <-done // Блокировка: ждем тру от очередного воркера
    }
}
