package main

import "log"

func main () {
	var myString string
	myString = "Green"

	log.Println("My string is set to", myString)
	
	changeUsingPointer(&myString)

	log.Println("After func call myString is set to", myString)
}

func changeUsingPointer (s *string) {
	log.Println("s is set to", s)

	newValue := "Red"
	*s = newValue
}