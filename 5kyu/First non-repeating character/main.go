// Write a function named first_non_repeating_letter that takes a string input, and returns the first character that is not repeated anywhere in the string.

// For example, if given the input 'stress', the function should return 't', since the letter t only occurs once in the string, and occurs first in the string.

// As an added challenge, upper- and lowercase letters are considered the same character, but the function should return the correct case for the initial letter.
// For example, the input 'sTreSS' should return 'T'.

// If a string contains all repeating characters, it should return an empty string ("") or None -- see sample tests.

package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(FirstNonRepeating("hello"))
	fmt.Println(FirstNonRepeating("moonmen"))
	fmt.Println(FirstNonRepeating("hello world, eh?"))
	fmt.Println(FirstNonRepeating("abba"))

}

func FirstNonRepeating(str string) string {
	var hash = map[string]int{}
	lowStr := strings.ToLower(str)

	for i := 0; i < len(lowStr); i++ {
		hash[lowStr[i:i+1]]++
	}

	for i := 0; i < len(lowStr); i++ {
		if hash[lowStr[i:i+1]] > 1 {
			continue
		}
		return str[i : i+1]
	}
	return ""
}
