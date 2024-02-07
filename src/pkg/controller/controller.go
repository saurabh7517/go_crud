package controller

import (
	"crud/src/pkg/service"
	"encoding/json"
	"fmt"
	"net/http"
)

var movies []service.Movie

func RegisterRestController(mux *http.ServeMux) {
	mux.HandleFunc("/movies", getAllMovies)
	mux.HandleFunc("/movie/id", processMovieById)
	mux.HandleFunc("/movie", createMovie)
}

type responseError struct {
	resWriter http.ResponseWriter
	errMsg    string
}

func (re responseError) writeError(errorMsg string) {
	re.resWriter.WriteHeader(http.StatusMethodNotAllowed)
	if errorMsg != "" {
		fmt.Fprint(re.resWriter, errorMsg)
	} else {
		fmt.Fprint(re.resWriter, re.errMsg)
	}
}

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		var movies []service.Movie = service.GetAllMovies()
		json.NewEncoder(w).Encode(movies)
		return
	}
	responseError{w, "Wrong http verb used"}.writeError("")
}

func processMovieById(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getMovieById(w, r)
	case http.MethodPut:
		updateMovie(w, r)
	case http.MethodDelete:
		deleteMovie(w, r)
	default:
		responseError{w, "Wrong http verb used"}.writeError("")
	}
}

func getMovieById(w http.ResponseWriter, r *http.Request) {

}

func createMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		return

	}
	responseError{w, "Wrong http verb used"}.writeError("")
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		return

	}
	responseError{w, "Wrong http verb used"}.writeError("")
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		return

	}
	responseError{w, "Wrong http verb used"}.writeError("")
}
