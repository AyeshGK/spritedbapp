package main

import (
	"log"
	"net/http"

	"github.com/AyeshGK/spritedbapp/app/api"
	"github.com/AyeshGK/spritedbapp/app/handlers"
	"github.com/AyeshGK/spritedbapp/app/sessions"
)

func main() {
	sessions.InitializeSession()

	http.HandleFunc("/", handlers.IndexHandler)

	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/api/login", api.LoginAPIHandler)
	http.HandleFunc("/api/logout", api.LogoutAPIHandler)

	http.HandleFunc("/dashboard", handlers.DashboardHandler)

	http.HandleFunc("/update", handlers.UpdateHandler)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
