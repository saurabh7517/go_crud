package service

import (
	"errors"
)

type Movie struct {
	Id       int       `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Response struct {
	Msg string `json:"msg"`
}

var movies []Movie

var unknownMovie Movie = Movie{-1, "garbage", "garbage", &Director{"garbage", "garbage"}}

func setMovies(newMovies []Movie) {
	movies = newMovies
}

func GenerateMovieList() {
	movies = append(movies, Movie{1, "123341", "The New Gig", &Director{"John", "Doe"}})
	movies = append(movies, Movie{2, "123342", "Wrecmeister Harmonies", &Director{"Bela", "Tar"}})
	movies = append(movies, Movie{3, "123343", "In Vanda's Room", &Director{"Pedro", "Costa"}})
	movies = append(movies, Movie{4, "123344", "The Intruder", &Director{"Claire", "Denis"}})
	movies = append(movies, Movie{5, "123345", "Morvern Callar", &Director{"Lynne", "Ramsay"}})
	movies = append(movies, Movie{6, "123346", "Possession", &Director{"Andrej", "Zulawski"}})
	movies = append(movies, Movie{7, "123347", "Earth", &Director{"Alaxender", "Dovzhenko"}})
	movies = append(movies, Movie{8, "123348", "Annie Hall", &Director{"Woody", "Allen"}})
	movies = append(movies, Movie{9, "123349", "Sullivan Travels", &Director{"Preston", "Sturges"}})
	movies = append(movies, Movie{10, "123350", "Pandora's Box", &Director{"G.W", "Pabst"}})
	movies = append(movies, Movie{11, "123351", "Born in Flames", &Director{"Lizzie", "Borden"}})
	movies = append(movies, Movie{12, "123352", "Grave of the Fireflies", &Director{"Isao", "Takahata"}})
}

func GetAllMovies() []Movie {
	if len(movies) > 0 {
		return movies
	} else {
		GenerateMovieList()
		return movies
	}
}

func GetMovieById(id int) (Movie, error) {
	for _, v := range movies {
		if id == v.Id {
			return v, nil
		}
	}
	return unknownMovie, errors.New("Not Found")
}

func AddNewMovie(newMovie Movie) Response {
	var primaryKey int = createNewPrimaryKey()
	newMovie.Id = primaryKey
	movies = append(movies, newMovie)
	return Response{"New movie added"}
}

func createNewPrimaryKey() int {
	var max int = -1
	for _, movie := range movies {
		if movie.Id > max {
			max = movie.Id
		}
	}
	return max + 1
}

func UpdateMovie(newMovie Movie) bool {
	var found bool = false
	for _, movie := range movies {
		if movie.Id == newMovie.Id {
			updateMovie(movie, newMovie)
			found = true
			break
		}
	}
	if !found {
		AddNewMovie(newMovie)
		return true
	} else {
		return false
	}
}

func updateMovie(oldMovie Movie, newMovie Movie) {
	oldMovie.Id = newMovie.Id
	oldMovie.Isbn = newMovie.Isbn
	oldMovie.Title = newMovie.Title
	oldMovie.Director.Firstname = newMovie.Director.Firstname
	oldMovie.Director.Lastname = newMovie.Director.Lastname
}

func RemoveMovie(id int) bool {
	var removedMovies []Movie = removeMovieById(id)
	setMovies(removedMovies)
	return true
}

func removeMovieById(id int) []Movie {
	var newMovies []Movie
	for _, movie := range movies {
		if movie.Id != id {
			newMovies = append(newMovies, movie)
		}
	}
	return newMovies
}
