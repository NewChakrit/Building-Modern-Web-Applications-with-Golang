package main

import (
	"fmt"
	"html/template"
	"net/http"
)


func renderTemplate (w http.ResponseWriter,html string ){
	parseTemplate, _ := template.ParseFiles("./templates/" + html)

	err:=parseTemplate.Execute(w, nil)
	if err!= nil{
		fmt.Println("Error parsing template", err)
		return
	}
}