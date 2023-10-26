// The function is not returning the correct values. Can you figure out why?

// Example (Input --> Output ):
// 3 --> "Earth"

package main

import "fmt"

func main() {

	fmt.Println(GetPlanetName(1))
	fmt.Println(GetPlanetName(5))
	fmt.Println(GetPlanetName(8))
	fmt.Println(GetPlanetName(2))

}

func GetPlanetName(ID int) string {

	switch ID {
	case 1:
		return "Mercury"
	case 2:
		return "Venus"
	case 3:
		return "Earth"
	case 4:
		return "Mars"
	case 5:
		return "Jupiter"
	case 6:
		return "Saturn"
	case 7:
		return "Uranus"
	case 8:
		return "Neptune"
	default:
		return "Pluto"
	}

}
