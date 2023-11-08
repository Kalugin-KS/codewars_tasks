// Write a function that accepts an integer n and a string s as parameters, and returns a string of s repeated exactly n times.

// Examples (input -> output)
// 6, "I"     -> "IIIIII"
// 5, "Hello" -> "HelloHelloHelloHelloHello"

package main

import "fmt"

func main() {

	fmt.Println(RepeatStr(4, "a"))
	fmt.Println(RepeatStr(3, "hello "))
	fmt.Println(RepeatStr(2, "abc"))

}

func RepeatStr(repetitions int, value string) string {
	var result string
	for i := 1; i <= repetitions; i++ {
		result += value
	}
	return result
}
