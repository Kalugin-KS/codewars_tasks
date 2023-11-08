// In this simple assignment you are given a number and have to make it negative. But maybe the number is already negative?

// Examples
// MakeNegative(1)    // return -1
// MakeNegative(-5)   // return -5
// MakeNegative(0)    // return 0

// The number can be negative already, in which case no change is required.
// Zero (0) is not checked for any specific sign. Negative zeros make no mathematical sense.

package main

import "fmt"

func main() {

	fmt.Println(MakeNegative(42))
	fmt.Println(MakeNegative(0))
	fmt.Println(MakeNegative(-4))

}

func MakeNegative(x int) int {
	if x > 0 {
		return -x
	} else {
		return x
	}
}
