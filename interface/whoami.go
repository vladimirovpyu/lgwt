package main

import "fmt"

func printType(whoami interface{}) {
	fmt.Printf("%T\n", whoami)
}

func main() {
	printType(42)          // выведет: int
	printType("im string") // выведет: string
	printType(true)        // выведет: bool
}
