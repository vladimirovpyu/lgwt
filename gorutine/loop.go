package main

import (
    "fmt"
    "sync"
    "time"
    )

func main() {
    wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
	    wg.Add(1)
		go func(v int) {
            defer fmt.Println("defer")
            defer time.Sleep(1 * time.Second)
            defer wg.Done()
			fmt.Println(v)
		}(i)
	}

    wg.Wait()
    //time.Sleep(1 * time.Second)
}
