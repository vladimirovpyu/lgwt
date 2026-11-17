package main

import (
    "net/http"
    "fmt"
    "sync"
    "io"
    "time"
)

func worker(ch <- chan string, wg *sync.WaitGroup) {
    defer wg.Done()
    client := &http.Client{
        Timeout: 10 * time.Second,
    }

    url := <-ch
    req, err := http.NewRequest("GET", url, nil)
    if (err != nil) {
       fmt.Println("Ошибка создания запроса:", err)
       return
    }

    resp, err := client.Do(req)
    if (err != nil) {
       fmt.Println("Ошибка выполнения запроса:", err)
       return
    }

    body, err := io.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err != nil {
        fmt.Println("Ошибка чтения ответа:", err)
        return
    }

    fmt.Printf("body size: %d\n", len(body));
}

func main(){
    ch := make(chan string, 2)
    ch <- "http://ya.ru"
    ch <- "http://google.ru"
    defer close(ch)

    var wg sync.WaitGroup
    wg.Add(2)

    for i:=0; i<2; i++ {
        go worker(ch, &wg)
    }

    for i:=0; i<2; i++ {
        wg.Wait()
    }
}
