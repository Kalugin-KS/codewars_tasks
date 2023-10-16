//Given a set of numbers, return the additive inverse of each. Each positive becomes negatives, and the negatives become positives.
//You can assume that all values are integers. Do not mutate the input array/list.

package main

import "fmt"

func main() {

	fmt.Println(Invert([]int{1, 2, 3, 4, 5}))
	fmt.Println(Invert([]int{1, -2, 3, -4, 5}))
}

func Invert(arr []int) []int {
	var result []int
	for _, val := range arr {
		result = append(result, -val)
	}
	return result
}
