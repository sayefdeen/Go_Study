package main

import (
	"fmt"
)

func main() {
	statePopulate := map[string]int{
		"Jordan": 12345678,
		"Egypt":  87654321,
		"Syria":  34567789,
	}
	// This will print 0, since Jorda is not available in the map
	// to prevent this to cause some confusion, we can use comma ok syntax
	// , ok
	// This will print 0 false, and we can check on the ok variable
	prop, ok := statePopulate["Jorda"]
	fmt.Println(prop, ok)
}

func usingMake() {
	statePopulation := make(map[string]int)
	a := make([]int, 3)
	fmt.Println(statePopulation)
	fmt.Println(a)
}
