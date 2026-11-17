package main

import (
    "fmt"
    )

func write() <- chan int{
    ch := make(chan int)

    go func() {
        defer close(ch)
        for i := range 2 {
            ch <- i + 1
        }
    }()

    return ch
}

func main() {
    ch := write()

    for {
        v, ok := <- ch
        if !ok {
            break
        }
        fmt.Println("v =", v)
    }
 }
