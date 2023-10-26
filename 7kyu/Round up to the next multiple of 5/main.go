// Given an integer as input, can you round it to the next (meaning, "greater than or equal") multiple of 5?

// Examples:
// input:    output:
// 0    ->   0
// 2    ->   5
// 3    ->   5
// 21   ->   25
// -2   ->   0
// -5   ->   -5

// Input may be any positive or negative integer (including 0).
// You can assume that all inputs are valid integers.

package main

import "fmt"

func main() {

	fmt.Println(RoundToNext5(6))
	fmt.Println(RoundToNext5(1))
	fmt.Println(RoundToNext5(3))
	fmt.Println(RoundToNext5(-8))
	fmt.Println(RoundToNext5(12))

}

func RoundToNext5(n int) int {
	switch {
	case n%5 > 0:
		return n + (5 - n%5)
	case n%5 < 0:
		return n - n%5
	default:
		return n
	}
}
