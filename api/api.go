package api

import (
	"log"
	"net/http"

	"github.com/AyeshGK/spritedbapp/app/sessions"
)

func LoginAPIHandler(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	log.Print("login credentials: ", username, password)

	if username == "admin" && password == "password" {
		log.Print("Authentication successful")
		sessions.SetSession(w, r, "authenticated", true)
		http.Redirect(w, r, "/dashboard", http.StatusFound)
		return
	}

	// Authentication failed
	http.Redirect(w, r, "/", http.StatusFound)
}

func LogoutAPIHandler(w http.ResponseWriter, r *http.Request) {
	sessions.ClearSession(w, r)
	http.Redirect(w, r, "/", http.StatusFound)
}
