package main

import "fmt"

func main() {
Loop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}

	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}

	array := []int{1, 2, 3}
	for _, v := range array {
		fmt.Println(v)
	}
}
