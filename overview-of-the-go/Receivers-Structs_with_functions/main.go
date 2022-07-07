package main

import "log"


type myStruck struct{
	Firstname string
}

func (m *myStruck) printFirstName () string {
	return m.Firstname
}

func main() {
	var myVar myStruck
	myVar.Firstname = "New"

	myVar2 := myStruck{
		Firstname: "Oum",
	}

	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.printFirstName())
}
