package main

import (
	"log"
	"time"
)

// var s = "seven"

// var firstName string
// var lastName string
// var phoneNumber string
// var age int
// var birthDate time.Time

type User struct {
	FirstName string
	LastName string 
	PhoneNumber string
	Age int
	BirthDate time.Time
}

func main () { 
	//  s2 := "six"

	//  s := "eight"
	// log.Println("s is", s)
	// log.Println("s2 is", s2)

	// saySomeThing("xxx")

	user:= User{ 
		FirstName: "New",
		LastName: "Chakrit",
		PhoneNumber: "555-5555-1212",
	}

	log.Println(user.FirstName, user.LastName, 
		user.BirthDate, user.PhoneNumber, user.Age)
}

// func saySomeThing (s3 string) (string, string) {
// 	log.Println("s from the saySomeThing func is", s)
// 	return s3, "World"
// }

