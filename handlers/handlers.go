package handlers

import (
	"html/template"
	"net/http"

	"github.com/AyeshGK/spritedbapp/app/sessions"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "login.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	if !sessions.IsAuthenticated(r) {
		http.Redirect(w, r, "/login", http.StatusFound) // Redirect to login page if not authenticated
		return
	}

	if err := templates.ExecuteTemplate(w, "dashboard.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	data := "This is updated content"
	if err := templates.ExecuteTemplate(w, "partial.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
