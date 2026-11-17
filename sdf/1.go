package main

/*
Дана строка, содержащая символы (){}.

Нужно определить, является ли входная строка сбалансированной с точки зрения скобок.

))
*/

import "fmt"

func isBalanced(s string) bool {

    if len(s) == 0 || len(s) % 2 == 1 {
        return false
    }

    stack := make([]byte, 0, len(s)/2)
    for i:=0; i < len(s); i++ {
        if (s[i] == '(' || s[i] == '[' || s[i] == '{') {
            stack = append(stack, s[i])
            continue
        }
        if s[i] == ')' && (len(stack) == 0 || stack[len(stack)-1] != '(') {
            return false;
        }
        if s[i] == ']' && (len(stack) == 0 || stack[len(stack)-1] != '[') {
            return false;
        }
        if s[i] == '}' && (len(stack) == 0 || stack[len(stack)-1] != '{') {
            return false;
        }
        stack = stack[:len(stack)-1]
    }

	return len(stack) == 0
}

func main() {
	fmt.Println(isBalanced("(())"))   // -> true
	fmt.Println(isBalanced("({}())")) // -> true
	fmt.Println(isBalanced("(()"))    // -> false
	fmt.Println(isBalanced("(()))"))  // -> false
	fmt.Println(isBalanced("({)}"))   // -> false
}
