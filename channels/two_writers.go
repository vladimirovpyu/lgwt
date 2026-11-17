package main

import (
    "fmt"
    "sync"
    )

func write() <- chan int{
    ch := make(chan int)
    wg := sync.WaitGroup{}
    wg.Add(2)

    go func() {
        defer wg.Done()
        for i := range 3 {
            ch <- i + 1
        }
    }()

    go func() {
        defer wg.Done()
        for i := range 3 {
            ch <- i + 11
        }
    }()

    go func() {
        wg.Wait()
        close(ch)
    }()

    return ch
}

func main() {
    ch := write()

    for v := range ch {
        fmt.Println("v =", v)
    }
 }
