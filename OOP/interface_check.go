package main

import "fmt"

type Worker interface {
	Work()
	Rest()
}

type Developer struct {
	Name string
}

// Реализуем только один метод из двух
func (d *Developer) Work() {
	fmt.Println(d.Name, "пишет код")
}

// 🛑 ЭТА СТРОКА ВЫЗОВЕТ ОШИБКУ КОМПИЛЯЦИИ!
// Мы явно просим компилятор проверить, является ли *Developer интерфейсом Worker.
var _ Worker = (*Developer)(nil)

func main() {
	dev := &Developer{Name: "Алексей"}
	dev.Work()
}

// ./interface_check.go:21:16: cannot use (*Developer)(nil) (value of type *Developer) as Worker value in variable
// declaration: *Developer does not implement Worker (missing method Rest)
