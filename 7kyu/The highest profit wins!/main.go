// Write a function that returns both the minimum and maximum number of the given list/array.

// Examples (Input --> Output)
// [1,2,3,4,5] --> [1,5]
// [2334454,5] --> [5,2334454]
// [1]         --> [1,1]
// Remarks
// All arrays or lists will always have at least one element, so you don't need to check the length.

package main

import "fmt"

func main() {

	fmt.Println(MinMax([]int{1, 2, 3, 4, 5}))
	fmt.Println(MinMax([]int{2334454, 5}))
	fmt.Println(MinMax([]int{11, 20, 3, 4, 5, 17, -6, 0}))

}

func MinMax(arr []int) [2]int {
	max, min := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	return [2]int{min, max}
}
