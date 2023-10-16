//Complete the solution so that it reverses the string passed into it.

package main

import "fmt"

func main() {

	fmt.Println(Solution("word"))
	fmt.Println(Solution("hello"))
	fmt.Println(Solution("h"))
}

func Solution(word string) string {
	var reverse []byte
	for i := len(word) - 1; i >= 0; i-- {
		reverse = append(reverse, word[i])
	}
	return string(reverse)
}
