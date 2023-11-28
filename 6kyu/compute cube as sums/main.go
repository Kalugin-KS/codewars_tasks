// In this Kata, you will be given a number n (where n >= 1) and your task is to find n consecutive odd numbers whose sum is exactly the cube of n.

// Mathematically:
// cube = n ** 3
// sum = a1 + a2 + a3 + ..... + an
// sum == cube
// a2 == a1 + 2, a3 == a2 + 2, .......

// For example:
// FindSummands(3) == []int{7,9,11} // because 7 + 9 + 11 = 27, the cube of 3. Note that there are only n = 3 elements in the array.
// Write function findSummands(n) which will return an array of n consecutive odd numbers with the sum equal to the cube of n (n*n*n).

package main

import "fmt"

func main() {

	fmt.Println(FindSummands(1))
	fmt.Println(FindSummands(3))
	fmt.Println(FindSummands(5))
	fmt.Println(FindSummands(10))
	fmt.Println(FindSummands(15))

}

func FindSummands(n int) []int {
	init := n*n - n + 1
	var arr []int

	for i := init; i <= (n-1)*2+init; i += 2 {
		arr = append(arr, i)
	}

	return arr
}
