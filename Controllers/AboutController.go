package Controllers

import (
	"html/template"
	"net/http"
)

func About(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./Resources/Views/About/About.html")
	tmpl.Execute(w, nil)

}
