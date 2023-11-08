/*
Given two arrays a and b write a function comp(a, b) (orcompSame(a, b)) that checks whether the two arrays have the "same" elements, with the same multiplicities (the multiplicity of a member is the number of times it appears). "Same" means, here, that the elements in b are the elements in a squared, regardless of the order.

Examples
Valid arrays
a = [121, 144, 19, 161, 19, 144, 19, 11]
b = [121, 14641, 20736, 361, 25921, 361, 20736, 361]
comp(a, b) returns true because in b 121 is the square of 11, 14641 is the square of 121, 20736 the square of 144, 361 the square of 19, 25921 the square of 161, and so on. It gets obvious if we write b's elements in terms of squares:

a = [121, 144, 19, 161, 19, 144, 19, 11]
b = [11*11, 121*121, 144*144, 19*19, 161*161, 19*19, 144*144, 19*19]
Invalid arrays
If, for example, we change the first number to something else, comp is not returning true anymore:

a = [121, 144, 19, 161, 19, 144, 19, 11]
b = [132, 14641, 20736, 361, 25921, 361, 20736, 361]
comp(a,b) returns false because in b 132 is not the square of any number of a.

a = [121, 144, 19, 161, 19, 144, 19, 11]
b = [121, 14641, 20736, 36100, 25921, 361, 20736, 361]
comp(a,b) returns false because in b 36100 is not the square of any number of a.

Remarks
a or b might be [] or {} (all languages except R, Shell).
a or b might be nil or null or None or nothing (except in C++, COBOL, Crystal, D, Dart, Elixir, Fortran, F#, Haskell, Nim, OCaml, Pascal, Perl, PowerShell, Prolog, PureScript, R, Racket, Rust, Shell, Swift).
If a or b are nil (or null or None, depending on the language), the problem doesn't make sense so return false. */

package main

import "fmt"

func main() {

	var a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
	var a2 = []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
	fmt.Println(Comp(a1, a2))

	a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
	a2 = []int{11 * 21, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
	fmt.Println(Comp(a1, a2))

	a1 = []int{}
	a2 = []int{25, 49}
	fmt.Println(Comp(a1, a2))

	a1 = []int{2, 3, 3}
	a2 = nil
	fmt.Println(Comp(a1, a2))

	a1 = []int{-121, -144, 19, -161, 19, -144, 19, -11}
	a2 = []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
	fmt.Println(Comp(a1, a2))

}

func Comp(array1 []int, array2 []int) bool {
	switch {
	case array1 == nil || array2 == nil:
		{
			return false
		}
	case len(array1) == 0 && len(array2) == 0:
		{
			return true
		}
	case len(array1) == 0 && len(array2) != 0:
		{
			return false
		}
	case len(array1) != 0 && len(array2) == 0:
		{
			return false
		}
	case len(array1) != len(array2):
		{
			return false
		}
	}

	count := 1
	for i := 0; i < len(array1); i++ {
		if count == 0 {
			return false
		}
		count = 0
		for j := 0; j < len(array2); j++ {
			if array1[i]*array1[i] == array2[j] {
				array2[j] = -1
				count++
				break
			}

		}
	}
	if count == 0 {
		return false
	}
	return true

}
