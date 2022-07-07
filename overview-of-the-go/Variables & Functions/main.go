package main

import "fmt"

// Variables & Functions

func main() {
	fmt.Println("Hello World")

	var whatToSay string
	var i int

	whatToSay = "Goodbye, crual world"
	fmt.Println(whatToSay)

	i = 7

	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThingThatWasSaid := saySomeThing()

	fmt.Println("The function return", whatWasSaid, theOtherThingThatWasSaid)
}

func saySomeThing () (string, string) {
	 return "Something", "Hi there"
}