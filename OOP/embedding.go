package main

import "fmt"

// Базовая структура
type Person struct {
	Name string
}

func (p *Person) Greet() {
	fmt.Printf("Привет, меня зовут %s\n", p.Name)
}

// Встраивание: Employee «вбирает в себя» Person
type Employee struct {
	Person    // Анонимное поле (встраивание)
	Company   string
}

func main() {
	// Инициализация требует явного указания встроенной структуры
	emp := Employee{
		Person:  Person{Name: "Алексей"},
		Company: "Google",
	}

	// 🚀 Магия встраивания: доступ к полям и методам напрямую!
	fmt.Println(emp.Name) // Выведет: Алексей (вместо emp.Person.Name)
	emp.Greet()           // Выведет: Привет, меня зовут Алексей
}
