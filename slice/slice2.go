package main

import "fmt"

func main() {
	sl := make([]string, 5, 7)
	f1(sl[1:2])
	fmt.Printf("%q\n", sl)

	sl2 := make([]string, 5, 7)
	f1(sl2)
	fmt.Printf("%q\n", sl2)
	fmt.Printf("%q\n", sl2[5:6])
}

func f1(sl []string) {
	sl = append(sl, "Hello World!")
}
