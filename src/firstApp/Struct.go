package main

import "fmt"

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

type Engineer struct {
	Number int
	Name   string
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "Saif",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
		},
	}

	fmt.Println(aDoctor)

	bDoctor := struct{ name string }{name: "Saif"}
	fmt.Println(bDoctor)
}
