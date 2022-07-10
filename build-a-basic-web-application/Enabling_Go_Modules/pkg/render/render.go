package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate (w http.ResponseWriter,html string ){
	parseTemplate, _ := template.ParseFiles("./templates/" + html, "./templates/base.layout.html")
	err:=parseTemplate.Execute(w, nil)
	if err!= nil{
		fmt.Println("Error parsing template", err)
		return
	}
}