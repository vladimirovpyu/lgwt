package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("Фоновая горутина:", i)
			runtime.Gosched() // Уступаем процессор другим горутинам
		}
	}()

	for i := 0; i < 3; i++ {
		fmt.Println("Главная горутина:", i)
		runtime.Gosched() // Даем шанс выполниться фоновой горутине
	}
}
