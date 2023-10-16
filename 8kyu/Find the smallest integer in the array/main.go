//Given an array of integers your solution should find the smallest integer.
//Supplied array will not be empty.

package main

import "fmt"

func main() {
	arr1 := []int{34, 15, 88, 2}
	arr2 := []int{32, -28, 108, -7}
	arr3 := []int{1000, -8, 111, 0, 8}

	fmt.Println(SmallestIntegerFinder(arr1))
	fmt.Println(SmallestIntegerFinder(arr2))
	fmt.Println(SmallestIntegerFinder(arr3))
}

func SmallestIntegerFinder(numbers []int) int {
	smallest := numbers[0]

	for i := range numbers {
		if numbers[i] < smallest {
			smallest = numbers[i]
		}
	}

	return smallest
}
