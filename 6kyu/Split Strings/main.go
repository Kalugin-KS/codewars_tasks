// Complete the solution so that it splits the string into pairs of two characters. If the string contains an odd number of characters then it should replace the missing second character of the final pair with an underscore ('_').

// Examples:
// 'abc' =>  ['ab', 'c_']
// 'abcdef' => ['ab', 'cd', 'ef']

package main

import "fmt"

func main() {
	fmt.Println(Solution("abc"))
	fmt.Println(Solution("dawsd"))
	fmt.Println(Solution("awsaws"))
}

func Solution(str string) []string {
	if len(str)%2 != 0 {
		str += "_"
	}
	res := []string{}
	for i := 0; i < len(str); i = i + 2 {
		res = append(res, str[i:i+2])
	}
	return res
}
