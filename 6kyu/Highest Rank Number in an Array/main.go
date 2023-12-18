// Complete the method which returns the number which is most frequent in the given input array. If there is a tie for most frequent number, return the largest number among them.
// Note: no empty arrays will be given.

// Examples
// [12, 10, 8, 12, 7, 6, 4, 10, 12]              -->  12
// [12, 10, 8, 12, 7, 6, 4, 10, 12, 10]          -->  12
// [12, 10, 8, 8, 3, 3, 3, 3, 2, 4, 10, 12, 10]  -->   3

package main

import "fmt"

func main() {

	fmt.Println(HighestRank([]int{12, 10, 8, 12, 7, 6, 4, 10, 12}))
	fmt.Println(HighestRank([]int{12, 10, 8, 12, 7, 6, 4, 10, 12, 10}))
	fmt.Println(HighestRank([]int{12, 10, 8, 8, 3, 3, 3, 3, 2, 4, 10, 12, 10}))
	fmt.Println(HighestRank([]int{2}))

}

func HighestRank(nums []int) int {
	var max, num int
	hash := map[int]int{}
	for _, val := range nums {

		if hash[val]++; hash[val] == max {
			max = hash[val]
			if num < val {
				num = val
			}
		}

		if hash[val] > max {
			max = hash[val]
			num = val
		}
	}
	return num
}
