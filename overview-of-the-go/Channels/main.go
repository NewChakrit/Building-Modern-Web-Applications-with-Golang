package main

import (
	"channels/helpers"
	"log"
)

const numPool = 1000

func Calculate(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
} 

func main () {
	intChan := make(chan int)
	defer close(intChan)

	go Calculate(intChan)

	num := <- intChan 
	log.Println(num)
}

