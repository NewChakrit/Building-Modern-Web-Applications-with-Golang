package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName string
}

func main () {
	// var myString string
	// var myInt int

	// myString = "Hi"
	// myInt = 11

	// mySecondString := "another string"  

	// log.Println(myString , mySecondString, myInt)

	// ---------------------------------

	myMap := make(map[string]string)

	myMap["dog"] = "Samson"
	
	myMap["other-dog"] = "Cassie"
	
	myMap["dog"] = "Fido"

	log.Println(myMap["dog"]) // Fido
	log.Println(myMap["other-dog"]) // Cassie

	// ----------------------------------
	
	myMap2 := make(map[string]int)
	
	myMap2["first"] = 1
	myMap2["second"] = 2
	
	log.Println(myMap2["first"], myMap2["second"]) // 1 2

	// ----------------------------------
	
	myMap3 := make(map[string]User)
	
	me := User{
		FirstName: "New",
		LastName: "Chakrit",
	}
	
	myMap3["me"] = me
	
	log.Println(myMap3["me"].FirstName) // New
	log.Println(myMap3["me"].LastName) // Chakrit
	log.Println(myMap3["me"]) // {New Chakrit}

	// myMap4 := make(map[string]interface{}) // interface{} --> can store every thing you want

	// ----------------- Slice (Array) -----------------

	// var mySlice[] string

	// mySlice = append(mySlice, "New") 
	// mySlice = append(mySlice, "John") 
	// mySlice = append(mySlice, "Mary") 
	
	// log.Println(mySlice) // [New John Mary]

	var mySlice[] int

	mySlice = append(mySlice, 2) 
	mySlice = append(mySlice, 1) 
	mySlice = append(mySlice, 3) 

	sort.Ints(mySlice)

	log.Println(mySlice) // [1 2 3]

	numbers := []int{1,2,3,4,5,6,7,8,9,10}

	log.Println(numbers) //[1 2 3 4 5 6 7 8 9 10]
	
	log.Println(numbers[0:2]) //[1 2]
	log.Println(numbers[6:8]) //[7 8]

	names := []string{"one", "seven", "fish", "cat"}

	log.Println(names) //[one seven fish cat]
}
