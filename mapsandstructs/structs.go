package mapsandstructs

import (
	"fmt"
)

type doctor struct {
	number     int
	actorName  string
	companions []string
}

//StructPlay1 creates a medical doctor struct and prints it
func StructPlay1() {
	aDoctor := doctor{
		number:    3,
		actorName: "Sander Furre",
		companions: []string{
			"Sandorini",
			"Furitius",
			"Martinio",
		},
	}

	fmt.Println(aDoctor)
}

type animal struct {
	name   string
	origin string
}

type bird struct {
	animal //Embedds animal, NOT inheritence
	//Has birdlike characteristics, but is not of tyupe animal
	speedKPH float32
	canFly   bool
}

//StructPlay2 provides structure embedding
func StructPlay2() {
	b := bird{
		animal:   animal{name: "Emu", origin: "Australia"},
		speedKPH: 48,
		canFly:   false,
	}

	//Look, we dont have to consider the underlying structure here!
	fmt.Println(b.name)

}
