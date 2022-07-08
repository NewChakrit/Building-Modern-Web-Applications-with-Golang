package main

import "log"

func main () {

	// ------------ Loop Int ------------

	// for i := 0; i<= 5 ; i++{
	// 	log.Println(i)
	// }

	// ------------ Loop Slice ------------

	animals:= []string{"cat","fish","dog","bird","cow"}

	// for i,animal:= range animals {
	// 	log.Println(i, animal)

	//  2022/07/08 17:35:50 0 cat
	//  2022/07/08 17:35:50 1 fish
	//  2022/07/08 17:35:50 2 dog
	//  2022/07/08 17:35:50 3 bird
	//  2022/07/08 17:35:50 4 cow
	// }

	for _,animal:= range animals {
		log.Println( animal)

		// 2022/07/08 17:35:50 cat
		// 2022/07/08 17:35:50 fish
		// 2022/07/08 17:35:50 dog
		// 2022/07/08 17:35:50 bird
		// 2022/07/08 17:35:50 cow
	}

	// ------------ Loop Map ------------

	animals2 := make(map[string]string)
	animals2["dog"] = "Fido"
	animals2["cat"] = "Fluffy"

	for animalType,animal:= range animals2 {
		log.Println(animalType, animal) 

		// 2022/07/08 17:41:12 dog Fido
		// 2022/07/08 17:41:12 cat Fluffy
	}

	// ------------ Loop String ------------

	// var firstLine = "Once upon a midnight dreary"

	// for i , l := range firstLine{
	// 	log.Println(i, ":", l)
	// }

	// ------------ Loop String ------------
	
	type User struct {
		FirstName string
		LastName string
		Email string
		Age int
	}

	var users [] User 
	users = append(users, User{"John", "Smith", "john@smith.com", 25,})
	users = append(users, User{"Mary", "Anderson", "mary@smith.com", 26,})
	users = append(users, User{"New", "Chakrit", "new@smith.com", 27,})
	users = append(users, User{"Ploy", "Oum", "ploy@smith.com", 28,})

	for _ , l := range users{
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}