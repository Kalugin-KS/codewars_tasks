// Count the number of Duplicates
// Write a function that will return the count of distinct case-insensitive alphabetic characters and numeric digits
// that occur more than once in the input string. The input string can be assumed to contain only alphabets (both uppercase and lowercase) and numeric digits.

// Example
// "abcde" -> 0 # no characters repeats more than once
// "aabbcde" -> 2 # 'a' and 'b'
// "aabBcde" -> 2 # 'a' occurs twice and 'b' twice (`b` and `B`)
// "indivisibility" -> 1 # 'i' occurs six times
// "Indivisibilities" -> 2 # 'i' occurs seven times and 's' occurs twice
// "aA11" -> 2 # 'a' and '1'
// "ABBA" -> 2 # 'A' and 'B' each occur twice

package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(duplicate_count("abcde"))
	fmt.Println(duplicate_count(""))
	fmt.Println(duplicate_count("indivisibility"))
	fmt.Println(duplicate_count("AABB"))
	fmt.Println(duplicate_count("adndg12bvfsa2edg1"))

}

func duplicate_count(s1 string) int {

	s1 = strings.ToLower(s1)
	count := 0
	hash := map[rune]int{}

	for _, val := range s1 {
		hash[val]++
		if hash[val] == 2 {
			count++
		}
	}
	return count
}
