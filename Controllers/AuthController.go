package Controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"example.com/events/database"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	id    int
	name  string
	phone string
}

type Session struct {
}

func Auth(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./Resources/Views/auth/auth.html")
	tmpl.Execute(w, nil)
}

func Register(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	phone := r.FormValue("phone")
	password := r.FormValue("password")
	hash, err := HashPassword(password)
	if err != nil {
		panic("err")
	}

	if phone == "" || password == "" {
		fmt.Println("Укажите логин и пароль")

	}
	database.InsertUser(name, phone, hash)
	getCookie(w, r)
	http.Redirect(w, r, "/index", http.StatusSeeOther)
}

func Authrorization(w http.ResponseWriter, r *http.Request) {
	phone := r.FormValue("phone")
	password := r.FormValue("password")
	hash, _ := HashPassword(password)
	fmt.Println(hash)
	database.GetUser(phone, hash)
	getCookie(w, r)
	http.Redirect(w, r, "/index", http.StatusSeeOther)
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

var Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func getCookie(w http.ResponseWriter, r *http.Request) {

	session, _ := Store.Get(r, "coockie-name")
	session.Values["authenticated"] = true
	session.Save(r, w)

}
