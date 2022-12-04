package Controllers

import (
	"html/template"
	"net/http"
)

func Main(w http.ResponseWriter, r *http.Request) {
	var user = User{}
	tmpl, _ := template.ParseFiles("./Resources/Views/Home/Home.html")
	tmpl.Execute(w, user)

}
