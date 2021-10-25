package main

import "fmt"

func main() {
	grades := [...]int{97, 85, 93}
	fmt.Println("Grades ", grades)
}

func something() {
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
}

func usingSlice() {
	a := []int{}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
}
