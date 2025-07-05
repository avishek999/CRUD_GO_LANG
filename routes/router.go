package routes

import (
	"GO_CRUD/handlers"
	"fmt"
	"net/http"
)

func RegisterRoutes() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to Go Movie API 🚀")
	})

	http.HandleFunc("/movies", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handlers.GetAllMovie(w, r)
		case "POST":
			handlers.CreateMovie(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/movies/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handlers.GetMovieByID(w, r)
		case "PUT":
			handlers.UpdateMovieById(w, r)
		case "DELETE":
			handlers.DeleteMovieById(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}
