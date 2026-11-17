package main

import "fmt"

func romanToInt(s string) int {

	array := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

    var result = 0
    for i := 0; i < len(s); i++ {
        result += array[string(s[i])]
        if (i > 0) {
            if array[string(s[i-1])] < array[string(s[i])] {
                result -= 2 * array[string(s[i-1])]
            }
        }
    }

    return result
}

func main() {
    fmt.Println(romanToInt("III"))
    fmt.Println(romanToInt("CDXLIV"))
}

// IV 4 IX 9
// XL 40 XC 90
// CD 400 CM 900

