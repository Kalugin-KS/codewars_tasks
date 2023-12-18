// Count the number of divisors of a positive integer n.

// Examples (input --> output)
// 4 --> 3 // we have 3 divisors - 1, 2 and 4
// 5 --> 2 // we have 2 divisors - 1 and 5
// 12 --> 6 // we have 6 divisors - 1, 2, 3, 4, 6 and 12
// 30 --> 8 // we have 8 divisors - 1, 2, 3, 5, 6, 10, 15 and 30

package main

import "fmt"

func main() {

	fmt.Println(Divisors(1))
	fmt.Println(Divisors(10))
	fmt.Println(Divisors(52))
	fmt.Println(Divisors(64))

}

func Divisors(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}

	return count
}
