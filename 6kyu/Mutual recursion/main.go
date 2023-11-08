// Mutual Recursion allows us to take the fun of regular recursion (where a function calls itself until a terminating condition) and apply it to multiple functions calling each other!

// Let's use the Hofstadter Female and Male sequences to demonstrate this technique. You'll want to create two functions F and M.

package main

import "fmt"

func main() {

	fmt.Println(F(0))
	fmt.Println(F(5))
	fmt.Println(F(10))
	fmt.Println(M(5))

}

func F(n int) int {
	if n == 0 {
		return 1
	}
	return n - M(F(n-1))
}

func M(n int) int {
	if n == 0 {
		return 0
	}
	return n - F(M(n-1))
}
