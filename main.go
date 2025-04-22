package main

import (
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
	"fmt"
)


func main() {
    r := chi.NewRouter()

    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    r.Handle("/src/*", http.StripPrefix("/src/", http.FileServer(http.Dir("src"))))

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "index.html")
    })

    http.ListenAndServe(":3000", r)
	fmt.Println("Server starting using Vite @ http://localhost:5173")
}
