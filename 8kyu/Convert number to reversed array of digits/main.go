// Given a random non-negative number, you have to return the digits of this number within an array in reverse order.

// Example(Input => Output):
// 35231 => [1,3,2,5,3]
// 0 => [0]

package main

import "fmt"

func main() {

	fmt.Println(Digitize(23582357))
	fmt.Println(Digitize(548413589698465))
	fmt.Println(Digitize(4741184625592156168))

}

func Digitize(n int) []int {
	var ret []int
	for {
		ret = append(ret, n%10)
		n /= 10
		if n == 0 {
			return ret
		}
	}
}
