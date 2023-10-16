// Write a function which reduces fractions to their simplest form! Fractions will be presented as an array/tuple (depending on the language) of strictly positive integers,
// and the reduced fraction must be returned as an array/tuple:

// input:   [numerator, denominator]
// output:  [reduced numerator, reduced denominator]
// example: [45, 120] --> [3, 8]
// All numerators and denominators will be positive integers.

package main

import "fmt"

func main() {

	fmt.Println(ReduceFraction([2]int{60, 20}))
	fmt.Println(ReduceFraction([2]int{80, 120}))
	fmt.Println(ReduceFraction([2]int{45, 120}))
	fmt.Println(ReduceFraction([2]int{1000, 1}))
	fmt.Println(ReduceFraction([2]int{1, 1}))
	fmt.Println(ReduceFraction([2]int{10956590, 13611876}))
}

func ReduceFraction(fraction [2]int) [2]int {

	result := fraction
	if result[0] == 1 || result[1] == 1 {
		return result
	}

	n := 2
	for {
		if result[0] == 1 || result[1] == 1 {
			break
		}
		if result[0]%n != 0 || result[1]%n != 0 {
			if result[0] == n {
				break
			}
			n++
			continue
		}

		result[0] = result[0] / n
		result[1] = result[1] / n
		n = 2
	}

	return result
}
