/* Write a program that will calculate the number of trailing zeros in a factorial of a given number.

N! = 1 * 2 * 3 *  ... * N

Be careful 1000! has 2568 digits...

For more info, see: http://mathworld.wolfram.com/Factorial.html

Examples
zeros(6) = 1
# 6! = 1 * 2 * 3 * 4 * 5 * 6 = 720 --> 1 trailing zero

zeros(12) = 2
# 12! = 479001600 --> 2 trailing zeros */

package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(Zeros(0))
	fmt.Println(Zeros(30))
	fmt.Println(Zeros(100))
	fmt.Println(Zeros(1000))

}

func Zeros(n int) int {
	var sum, i int
	for {
		i++
		res := int(float64(n) / math.Pow(5, float64(i)))
		if res == 0 {
			break
		}
		sum += res
	}
	return sum
}
