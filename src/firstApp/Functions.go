package main

import "fmt"

func main() {
	a := "Strart"
	defer fmt.Println(a)
	a = "End"
}
