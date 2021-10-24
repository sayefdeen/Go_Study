package main

import (
	"fmt"
)

// Variables at the package level

var i int = 42

var (
	something     string = "something"
	somethingelse string = "somethingelse"
)

func main() {
	i := 42

	fmt.Println(i)
	method()
}

func method() {
	s := "This is a string"
	b := []byte(s)
	d := string(b)
	fmt.Println(b)
	fmt.Println(d)
}

// Define a variable at compile time

func compileTime() {
	//  This will works since a will be created at compile time, the compiler will see when it is used and will define its type according to that
	const a = 42
	var b int16 = 27
	fmt.Println(a + b)
}
