package main

import "fmt"

func main() {
	sl := make([]string, 5, 7)
	f1(sl) // append записал строку на индекс 5

	// Достаем значение "через capacity":
	// Расширяем срез до его максимально возможной емкости (до 7)
	fullSl := sl[:cap(sl)]

	// Теперь мы можем безопасно напечатать весь базовый массив
	// или обратиться к конкретному индексу 5
	fmt.Printf("Весь массив: %q\n", fullSl)
	fmt.Printf("Конкретное значение: %q\n", fullSl[5])
}

func f1(sl []string) {
	sl = append(sl, "Hello World!")
}
