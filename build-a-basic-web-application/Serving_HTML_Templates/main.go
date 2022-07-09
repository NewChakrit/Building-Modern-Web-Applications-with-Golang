package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8080"

// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}

// About is the handler for the about page
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html")

}

func renderTemplate (w http.ResponseWriter,html string ){
	parseTemplate, _ := template.ParseFiles("./templates/" + html)

	err:=parseTemplate.Execute(w, nil)
	if err!= nil{
		fmt.Println("Error parsing template", err)
		return
	}
}

// main is the main function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

