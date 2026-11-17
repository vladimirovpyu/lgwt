package main

import (
    "fmt"
    "time"
    "math/rand"
)

func longTimeFunction(val int) int {
    time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
    return val * 2
}

func predictableTimeWork() {
    ch := make(chan struct{})

    go func() {
        longTimeFunction(1)
        close(ch)
    }()

    select {
        case <- ch:
            fmt.Println("chan")
        case <- time.After(3 * time.Second):
            fmt.Println("3 sec")
    }

}

func main() {
    predictableTimeWork()
}
