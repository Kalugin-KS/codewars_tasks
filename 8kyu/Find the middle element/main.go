// As a part of this Kata, you need to create a function that when provided with a triplet, returns the index of the numerical element that lies between the other two elements.
// The input to the function will be an array of three distinct numbers (Haskell: a tuple).

// For example:
// gimme([2, 3, 1]) => 0
// 2 is the number that fits between 1 and 3 and the index of 2 in the input array is 0.

package main

import "fmt"

func main() {

	fmt.Println(Gimme([3]int{2, 3, 1}))
	fmt.Println(Gimme([3]int{5, 10, 14}))
	fmt.Println(Gimme([3]int{15, 10, 14}))
	fmt.Println(Gimme([3]int{-15, -10, 14}))

}

func Gimme(array [3]int) int {
	res := 0

	if (array[1] > array[0] && array[1] < array[2]) || (array[1] < array[0] && array[1] > array[2]) {
		res = 1
	}
	if (array[2] > array[1] && array[2] < array[0]) || (array[2] < array[1] && array[2] > array[0]) {
		res = 2
	}
	return res
}
