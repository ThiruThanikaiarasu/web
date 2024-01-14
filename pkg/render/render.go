package render 

import (
	"fmt"
	"net/http"
	"html/template"
)

// RenderedTemplate renders template using html/template 
func RenderedTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.gohtml")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template ", err)
		return 
	}
} 