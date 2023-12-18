// In this Kata, you will be given a string and your task will be to return a list of ints detailing the
// count of uppercase letters, lowercase, numbers and special characters, as follows.

// Solve("*'&ABCDabcde12345") = [4,5,5,3].
// --the order is: uppercase letters, lowercase, numbers and special characters.

package main

import "fmt"

func main() {
	fmt.Println(Solve("Codewars@codewars123.com"))
	fmt.Println(Solve("P*K4%>mQUDaG$h=cx2?.Czt7!Zn16p@5H"))
	fmt.Println(Solve("RYT'>s&gO-.CM9AKeH?,5317tWGpS<*x2ukXZD"))
}

func Solve(s string) []int {
	res := make([]int, 4)
	for i := range s {
		switch {
		case s[i] >= 65 && s[i] <= 90:
			res[0]++
		case s[i] >= 97 && s[i] <= 122:
			res[1]++
		case s[i] >= 48 && s[i] <= 57:
			res[2]++
		default:
			res[3]++
		}
	}
	return res

}
