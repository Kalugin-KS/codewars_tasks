// Write a function that takes an array of strings as an argument and returns a sorted array containing the same strings, ordered from shortest to longest.
// For example, if this array were passed as an argument:
// ["Telescopes", "Glasses", "Eyes", "Monocles"]

// Your function would return the following array:
// ["Eyes", "Glasses", "Monocles", "Telescopes"]

// All of the strings in the array passed to your function will be different lengths, so you will not have to decide how to order multiple strings of the same length.

package main

import "fmt"

func main() {

	fmt.Println(SortByLength([]string{"beg", "life", "i", "to"}))
	fmt.Println(SortByLength([]string{"", "moderately", "brains", "pizza"}))
	fmt.Println(SortByLength([]string{"longer", "longest", "short"}))

}

func SortByLength(arr []string) []string {

	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if len(arr[j]) < len(arr[i]) {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	return arr

}
