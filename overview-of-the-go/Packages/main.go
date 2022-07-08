package main

import "log"

func main () {
	log.Println("Hello")

	var myVar helpers.SomeType
	myVar.TypeName = "New"
	log.Println(myVar.TypeName)
}