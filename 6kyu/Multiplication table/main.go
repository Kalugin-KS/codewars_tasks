// Your task, is to create N×N multiplication table, of size provided in parameter.

// For example, when given size is 3:

// 1 2 3
// 2 4 6
// 3 6 9

// For the given example, the return value should be:
// [[1,2,3],[2,4,6],[3,6,9]]

package main

import "fmt"

func main() {
	fmt.Println(MultiplicationTable(1))
	fmt.Println(MultiplicationTable(5))
	fmt.Println(MultiplicationTable(10))
}

func MultiplicationTable(size int) [][]int {
	arr := make([][]int, size)
	for i := 0; i < size; i++ {
		arr[i] = make([]int, size)
		for j := 0; j < size; j++ {
			arr[i][j] = (i + 1) * (j + 1)
		}
	}
	return arr
}
