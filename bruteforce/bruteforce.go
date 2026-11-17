package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"time"
)

// Наш алфавит (переменная alphabet из условия задачи)
var alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h"}

// hashPassword принимает строку и возвращает её MD5 хэш в виде слайса байт
func hashPassword(password string) []byte {
	hasher := md5.New()
	hasher.Write([]byte(password))
	return hasher.Sum(nil)
}

// RecoverPassword — ваша функция перебора, которую мы разбирали
func RecoverPassword(targetHash []byte) string {
	if len(alphabet) == 0 {
		return ""
	}

	// Перебираем длины пароля от 1 до 8 символов
	for length := 1; length <= 8; length++ {
		indices := make([]int, length)

		for {
			// Собираем текущий вариант пароля
			var buf bytes.Buffer
			for _, idx := range indices {
				buf.WriteString(alphabet[idx])
			}
			currentPassword := buf.String()

			// Сравниваем хэши
			if bytes.Equal(hashPassword(currentPassword), targetHash) {
				return currentPassword
			}

// 0 0 0
// 0 0 1
// 0 0 2
// 0 1 0
// 0 1 1
// ...
// 2 2 2

			// Логика инкремента индексов
			incremented := false
			for i := length - 1; i >= 0; i-- {
				if indices[i] < len(alphabet)-1 {
					indices[i]++
					incremented = true
					break
				}
				indices[i] = 0
			}

			if !incremented {
				break
			}
		}
	}
	return ""
}

func main() {
	// 1. Задаем секретный пароль (используем только символы из нашего алфавита!)
	secretPassword := "hgfedc"
	fmt.Printf("Исходный пароль: %s\n", secretPassword)

	// 2. Генерируем его хэш (эмулируем «базу данных» из условия)
	targetHash := hashPassword(secretPassword)
	fmt.Printf("Хэш в БД (MD5): %x\n", targetHash)

	fmt.Println("----------------------------------------")
	fmt.Println("Запуск восстановления пароля...")

	startTime := time.Now()
	// 3. Запускаем алгоритм восстановления
	restoredPassword := RecoverPassword(targetHash)
	duration := time.Since(startTime)

	// 4. Проверяем результат
	if restoredPassword == secretPassword {
		fmt.Printf("Успех! Пароль успешно восстановлен: %s\n", restoredPassword)
		fmt.Printf("Время перебора: %v\n", duration)
	} else {
		fmt.Println("Ошибка: Пароль не найден или алгоритм работает неверно.")
	}
}
