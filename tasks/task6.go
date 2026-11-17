package main

import (
	"fmt"
	"sync"
)

var globalMap = map[string][]int{
	"test":  make([]int, 0),
	"test2": make([]int, 0),
	"test3": make([]int, 0),
}
var a = 0

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex // мьютекс для защиты общих данных (a и globalMap)

	wg.Add(3)

	go func() {
		defer wg.Done() // гарантирует уменьшение счетчика только при выходе из горутины
		mu.Lock()
		defer mu.Unlock()
		a = 10
		globalMap["test"] = append(globalMap["test"], a)
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
		a = 11
		globalMap["test2"] = append(globalMap["test2"], a)
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
		a = 12
		globalMap["test3"] = append(globalMap["test3"], a)
	}()

	wg.Wait()
	fmt.Printf("%v\n", globalMap)
	fmt.Printf("%d\n", a)
}
