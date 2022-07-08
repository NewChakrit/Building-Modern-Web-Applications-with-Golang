package main

import "log"

func main () {
	var isTrue bool
	isTrue = true

	if isTrue == true {
		log.Println("isTrue is", isTrue)
		} else {
		log.Println("isTrue is", isTrue)
	}

	// ---------------------------------

	cat:="car2"

	if cat == "cat" {
		log.Println("Cat is cat")
	} else {
		log.Println("Cat is not cat")
	}

	// ---------------------------------

	myNum := 100
	isTrue2 := false

	if myNum > 99 && !isTrue2 {
		log.Println("myNum is greater than 99 and isTrue is set to true")

	} else if myNum < 100 && isTrue2 {
		log.Println("1")

	} else if myNum == 101 || isTrue2 {
	log.Println("2")
	
	} else if myNum > 100 && isTrue2 == false {
	log.Println("3")
	}

	// ----------------- Switch Case ----------------

	myVar := ""

	switch myVar {
	case "cat":
		log.Println("myVar is set to cat")
	case "dog":
		log.Println("myVar is set to dog")
	case "fish":
		log.Println("myVar is set to fish")
		default :
		 log.Println("myVar is somthing else")
	}
}