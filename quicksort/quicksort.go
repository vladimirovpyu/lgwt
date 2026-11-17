package main

import "fmt"

func quicksort(array []int) int {
    len := len(array)
    if len <=1 {
    	return 0
    }
    k := 0
    iterations := 0

    for i := 1; i < len; i++ {
        if array[i] < array[k] {
	        num := array[i]
            for j := i; j > k; j-- {
	            array[j] = array[j-1] // это не совсем правильно
	            iterations++
            }
            array[k] = num
            k = k + 1
        }
    }

    fmt.Println(array)

    if (k > 2) {
        iterations += quicksort(array[:k])
    }
    if (k < len - 2) {
        iterations += quicksort(array[k+1:])
    }

    return iterations
}

func main() {
    //array := []int{8, 16, 10, 2, 5, 3, 1, 4, 8, 16, 10, 2, 5, 3, 1, 4, 8, 16, 10, 2, 5, 3, 1, 4}
    array := []int{99, 16, 10, 2, 5, 3, 1, 4, 8, 16, 10, 2, 5, 3, 1, 4, 8, 16, 10, 2, 5, 3, 1, 4}


    iterations := quicksort(array[:])

    fmt.Println(array)
    fmt.Println(iterations)
}
// это не quicksort, скорее гибридный алгоритм