package main

import "fmt"

func main() {
	statePopulate := map[string]int{
		"Jordan": 12345678,
		"Egypt":  87654321,
		"Syria":  34567789,
	}

	if pop, ok := statePopulate["Jordan"]; ok {
		fmt.Println(pop)
	}
}

func switchStm() int {
	switch 3 {
	case 1, 2, 4:
		return 1

	case 3:
		return 3

	}
	return 0
}
