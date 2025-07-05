package handlers

import (
	"GO_CRUD/data"
	"GO_CRUD/models"
	"encoding/json"
	"net/http"
	"strings"
)




func GetAllMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/Json")
	json.NewEncoder(w).Encode(data.Movies)

}

func GetMovieByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/movies/")
	for _, movie := range data.Movies {
		if movie.ID == id {
			w.Header().Set("Content-Type", "application-json")
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {

	var newMovie models.Movie
	err := json.NewDecoder(r.Body).Decode((&newMovie))
	if err != nil {
		http.Error(w, "Invalid Body", http.StatusBadRequest)
		return
	}

	data.Movies = append(data.Movies, newMovie)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newMovie)
}

func DeleteMovieById(w http.ResponseWriter, r *http.Request) {

	id := strings.TrimPrefix(r.URL.Path, "/movies/")

	for i, movie := range data.Movies {
		if movie.ID == id {
			data.Movies = append(data.Movies[:i], data.Movies[i+1:]...)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Movie deleted"))
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)
}

func UpdateMovieById(w http.ResponseWriter, r *http.Request) {

	id := strings.TrimPrefix(r.URL.Path, "/movies/")

	var updated models.Movie

	err := json.NewDecoder(r.Body).Decode(&updated)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for i, movie := range data.Movies {
		if movie.ID == id {
			updated.ID = id
			data.Movies[i] = updated

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updated)
			return
		}

	}
	http.Error(w, "Movie not found", http.StatusNotFound)
}
