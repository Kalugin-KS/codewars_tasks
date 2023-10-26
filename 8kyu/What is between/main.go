// Complete the function that takes two integers (a, b, where a < b) and return an array of all integers between the input parameters, including them.

// For example:
// a = 1
// b = 4
// --> [1, 2, 3, 4]

package main

import "fmt"

func main() {

	fmt.Println(Between(1, 4))
	fmt.Println(Between(-4, 4))
	fmt.Println(Between(5, 50))
	fmt.Println(Between(11, 22))
}

func Between(a, b int) []int {
	var arr []int

	for i := a; i <= b; i++ {
		arr = append(arr, i)
	}

	return arr
}
