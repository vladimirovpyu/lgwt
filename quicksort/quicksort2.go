package main

import "fmt"

func quicksort(array []int) int {
    len := len(array)
    if len <=1 {
    	return 0
    }
    k := len-1
    pivot := array[k]
    i := 0
    iteration := 0

    fmt.Println("Before:")
    fmt.Println(array)

    for j := 0; j < len-1; j++ {
        if array[j] <= pivot {
	        array[i], array[j] = array[j], array[i]
	        iteration++
	        i++
        }
    }
    array[i], array[k] = array[k], array[i]

    fmt.Println("After:")
    fmt.Println(array)

    fmt.Printf("I = %d\n", i)

    if (i > 1) {
        iteration += quicksort(array[:i])
    }
    if (i + 2 < len) {
        iteration += quicksort(array[i:])
    }

    return iteration
}

func main() {
    array1 := []int{3, 2, 1}
    quicksort(array1[:])
    fmt.Println(array1)
    array2 := []int{99, 16, 10, 2, 5, 3, 1, 4, 8, 16, 10, 2, 5, 3, 1, 4, 8, 16, 10, 2, 5, 3, 1, 4}
    iterations := quicksort(array2[:])
    fmt.Println(array2)
    fmt.Println(iterations)
}
