// In this little assignment you are given a string of space separated numbers, and have to return the highest and lowest number.
// Output string must be two numbers separated by a single space, and highest number is first.
// Examples:
// HighAndLow("1 2 3 4 5")  // return "5 1"
// HighAndLow("1 2 -3 4 5") // return "5 -3"
// HighAndLow("1 9 3 4 -5") // return "9 -5"

package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(HighAndLow("7 0 -8 2 5 -11 1"))
	fmt.Println(HighAndLow("10 2 -1 -20"))
	fmt.Println(HighAndLow("42"))

}

func HighAndLow(in string) string {
	var arr []int
	var str string
	for i := 0; i < len(in); i++ {

		if in[i:i+1] != " " {
			str += in[i : i+1]

			if i+1 == len(in) {
				val, _ := strconv.Atoi(str)
				arr = append(arr, val)
			}

		} else {
			val, _ := strconv.Atoi(str)
			arr = append(arr, val)
			str = ""
		}

	}

	min, max := arr[0], arr[0]
	for i := range arr {
		switch {
		case min > arr[i]:
			min = arr[i]
		case max < arr[i]:
			max = arr[i]
		}
	}

	return strconv.Itoa(max) + " " + strconv.Itoa(min)
}
