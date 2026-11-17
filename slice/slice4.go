package main

import "fmt"

func a() {
	x := []int{}
	x = append(x, 0) // x len = 1 cap = 1 [0]
	x = append(x, 1) // x len = 2 cap = 2 [0 1]
	x = append(x, 2) // x len = 3 cap = 4 [0 1 2 _]
	y := append(x, 3) // x len = 3 cap = 4, результат [0 1 2 3]
	z := append(x, 4) // x len = 3 cap = 4, результат [0 1 2 4]
	fmt.Println(x, y, z) // т.к. y и z смотрят на один и тот же массив, результат [0 1 2 4] [0 1 2 4]

	y = append(x, 3) // x len = 3 cap = 4, результат [0 1 2 3]
	z = append(x, 4, 5) // x len = 3 cap = 4, z новый массив len=5, cap=8, результат [0 1 2 4 5]
	fmt.Println(x, y, z) // y и z смотрят на разные массивы, результат [0 1 2 3] [0 1 2 4 5]

}

func main() {
	a()
}
