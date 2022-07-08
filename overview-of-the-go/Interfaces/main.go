package main

import "fmt"

type Animal interface{
	Says() string
	NumberOfLegs() int
}

type Dog struct{
	Name string
	Breed string
}

type Gorilla struct {
	Name string
	Color string
	NumberOfTeeth int 
}
func main () {
	dog:= Dog {
		Name: "Samson",
		Breed: "German Shaphered",
	}

	PrintInfo(&dog)

	gorilla := Gorilla {
		Name: "Tim",
		Color: "gray",
		NumberOfTeeth: 33,
	}

	PrintInfo(&gorilla)
}

func PrintInfo(a Animal){
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")

	// This animal says Woof and has 4 legs
	// This animal says Ugh and has 2 legs
}

// ------------- Dog --------------

func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int{
	return 4
}

// ------------- Gorilla --------------

func (g *Gorilla) Says() string {
	return "Ugh"
}

func (g *Gorilla) NumberOfLegs() int{
	return 2
}