// Create a function with two arguments that will return an array of the first n multiples of x.
// Assume both the given number and the number of times to count will be positive numbers greater than 0.

// Examples:
// countBy(1,10)  should return  {1,2,3,4,5,6,7,8,9,10}
// countBy(2,5)  should return {2,4,6,8,10}

package main

import "fmt"

func main() {

	fmt.Println(CountBy(1, 5))
	fmt.Println(CountBy(2, 5))
	fmt.Println(CountBy(3, 5))
	fmt.Println(CountBy(50, 5))
	fmt.Println(CountBy(4, 7))

}

func CountBy(x, n int) []int {
	result := make([]int, n)
	for i := range result {
		result[i] = x * (i + 1)
	}
	return result
}
