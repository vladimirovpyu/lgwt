package main

import "fmt"

type Foo struct{}

func (f *Foo) A() {}
func (f *Foo) B() {}
func (f *Foo) C() {}

type AB interface {
	A()
	B()
}

type BC interface {
	B()
	C()
}

func main() {
	var f AB = &Foo{}

	// Этот type assertion сработает успешно, так как *Foo реализует интерфейс BC
	y := f.(BC)

	y.A() // <- ОШИБКА КОМПИЛЯЦИИ: y.A undefined (type BC has no field or method A)

	_ = y

	fmt.Println("Код успешно скомпилирован и выполнен!")
}
