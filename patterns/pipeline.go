package main

import (
    "fmt"
    "time"
    )

func writer() <- chan int {
    out := make(chan int)

    go func(){
        defer close(out)
        for i:=1; i<=10; i++ {
            out <- i
        }
    }()

    return out
}

func doubler(in <- chan int) <- chan int {
    out := make(chan int)

    go func(){
        defer close(out)
        for i := range in {
            out <- i * 2
            time.Sleep(500 * time.Millisecond)
        }
    }()

    return out
}

func reader(in <- chan int) {
    for i := range in {
        fmt.Println(i)
    }
}

func main() {
    reader(doubler(writer()))
}
