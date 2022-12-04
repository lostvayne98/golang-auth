package Routes

import (
	"net/http"

	"example.com/events/Controllers"
)

func Main() {
	http.HandleFunc("/login", Controllers.Auth)
	http.HandleFunc("/index", middleware(Controllers.Main))
	http.HandleFunc("/about", middleware(Controllers.About))
	http.HandleFunc("/registration", Controllers.Register)
	http.HandleFunc("/auth", Controllers.Authrorization)
	http.ListenAndServe(":3333", nil)

}

func middleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := Controllers.Store.Get(r, "cookie-name")

		// Check if user is authenticated
		if session.Values["authenticated"] == false {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next(w, r)
	}
}
