// Define a function that takes in two non-negative integers A and B and returns the last decimal digit of A^B. Note that A and B may be very large!
// For example, the last decimal digit of 9^7 is 9 since 9^7 = 4782969. The last decimal digit of (2^200)^(2^300) which has over 10^92 decimal digits, is 6. Also, please take 0^0=1

// Examples
// lastDigit 4 1             `shouldBe` 4
// lastDigit 4 2             `shouldBe` 6
// lastDigit 9 7             `shouldBe` 9
// lastDigit 10 (10^10)      `shouldBe` 0
// lastDigit (2^200) (2^300) `shouldBe` 6

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(LastDigit("4", "1"))
	fmt.Println(LastDigit("9", "7"))
	fmt.Println(LastDigit("10", "10000000000"))
	fmt.Println(LastDigit("3715290469715693021198967285016729344580685479654510946723", "68819615221552997273737174557165657483427362207517952651"))
}

func LastDigit(n1, n2 string) int {
	num1, _ := strconv.Atoi(n1[len(n1)-1:])
	num2, _ := strconv.Atoi(n2[len(n2)-1:])
	if len(n2) > 1 {
		num2, _ = strconv.Atoi(n2[len(n2)-2:])
	}
	switch num2 % 4 {
	case 1:
		return num1
	case 2:
		return (num1 * num1) % 10
	case 3:
		return (num1 * num1 * num1) % 10
	default:
		if num1%2 == 0 {
			if num1 == 0 {
				return 0
			}
			return 6
		}
		if num1 == 5 {
			return 5
		}
		return 1
	}
}
