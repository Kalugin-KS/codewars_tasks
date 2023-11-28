// The number 89 is the first integer with more than one digit that fulfills the property partially introduced in the title of this kata.
// Because this sum gives the same number: 89 = 8^1 + 9^2
// See this property again: 135 = 1^1 + 3^2 + 5^3

// We need a function to collect these numbers, that may receive two integers a, b that defines the range [a,b] (inclusive) and outputs
// a list of the sorted numbers in the range that fulfills the property described above.

// Examples
// Let's see some cases (input -> output):

// 1, 10  --> [1, 2, 3, 4, 5, 6, 7, 8, 9]
// 1, 100 --> [1, 2, 3, 4, 5, 6, 7, 8, 9, 89]

package main

import "fmt"

func main() {

	fmt.Println(SumDigPow(1, 10))
	fmt.Println(SumDigPow(100, 500))
	fmt.Println(SumDigPow(1, 100))
	fmt.Println(SumDigPow(50, 150))
	fmt.Println(SumDigPow(1000, 10000))

}

func SumDigPow(a, b uint64) []uint64 {
	result := make([]uint64, 0)
	for k := a; k <= b; k++ {
		num := k
		arr := make([]uint64, 0)
		for num > 0 {
			arr = append(arr, num%10)
			num = (num - num%10) / 10
		}
		var sum uint64
		for i, val := range arr {
			var multi uint64 = 1
			for j := len(arr) - i; j > 0; j-- {
				multi *= val
			}
			sum += multi
		}
		if sum == k {
			result = append(result, sum)
		}
	}
	return result
}
