package controller

import (
	"crud/src/pkg/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var movies []service.Movie

func RegisterRestController(mux *http.ServeMux) {
	mux.HandleFunc("/movies", getAllMovies)
	mux.HandleFunc("/movie", processMovieById)
}

type responseError struct {
	resWriter http.ResponseWriter
	errMsg    string
}

type responseMsg struct {
	msg  string
	code int
}

func writeError(re http.ResponseWriter, errorMsg string) {
	re.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(re).Encode(responseMsg{errorMsg, http.StatusMethodNotAllowed})
}

func writeResponse[T any](re http.ResponseWriter, movies T) {
	re.Header().Set("Content-Type", "application/json")
	json.NewEncoder(re).Encode(movies)
}

func readRequest(r *http.Request) (service.Movie, error) {
	var movie service.Movie
	var err error = json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		return movie, err
	}
	return movie, nil
}

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var movies []service.Movie = service.GetAllMovies()
		writeResponse(w, movies)
		return
	}
	writeError("Wrong http verb used")
}

func processMovieById(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getMovieById(w, r)
	case http.MethodPut:
		updateMovie(w, r)
	case http.MethodDelete:
		deleteMovie(w, r)
	case http.MethodPost:
		createMovie(w, r)
	default:
		writeError("Wrong http verb used")
	}
}

func getMovieById(w http.ResponseWriter, r *http.Request) {
	var movieIdString string = r.URL.Query().Get("id")
	movieId, err := strconv.ParseInt(movieIdString, 10, 32)
	if err != nil {
		fmt.Fprint(w, "Integer value not present as last parameter")
	}
	movie, err := service.GetMovieById(int(movieId))
	if err != nil {
		fmt.Fprint(w, movieId, " :: not present in database")
	}
	writeResponse(w, movie)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	movieDto, err := readRequest(r)
	if err != nil {
		responseError{w, "Error parsing body"}.writeError("")
		return
	}
	var response service.Response = service.AddNewMovie(movieDto)
	writeResponse(w, response)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	movieDto, err := readRequest(r)
	if err != nil {
		responseError{w, "Error parsing body"}.writeError("")
		return
	}
	var response service.Response = service.UpdateMovie(movieDto)
	writeResponse(w, response)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	var movieIdString string = r.URL.Query().Get("id")
	movieId, err := strconv.ParseInt(movieIdString, 10, 32)
	if err != nil {
		fmt.Fprint(w, "Integer value not present as last parameter")
		return
	}
	movieToDelete, err := service.GetMovieById(int(movieId))
	if err != nil {
		errorMsg := fmt.Sprintf("Movie with id %v not presnet in database", movieId)
		fmt.Fprint(w, errorMsg)
		return
	}
	var response service.Response = service.RemoveMovie(movieToDelete.Id)
	writeResponse(w, response)
}
