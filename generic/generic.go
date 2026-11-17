package main

import (
	"cmp"
	"fmt"
)

// [T cmp.Ordered] — это параметр типа с ограничением (constraint).
// cmp.Ordered означает, что T может быть любым типом, который поддерживает < и > (числа, строки).
func Min[T cmp.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Компилятор сам выводит тип T из переданных аргументов:
	fmt.Println(Min(10, 20))         // T = int, вернет int (10)
	fmt.Println(Min(3.14, 2.71))     // T = float64, вернет float64 (2.71)
	fmt.Println(Min("apple", "pear")) // T = string, вернет string ("apple")

	// Min(10, "pear") -> ОШИБКА КОМПИЛЯЦИИ! Смешивать разные типы нельзя.
}
