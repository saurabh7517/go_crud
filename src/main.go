package main

import (
	controller "crud/src/pkg/controller"
	"crud/src/pkg/service"
	"net/http"
)

func main() {
	var mux *http.ServeMux = http.NewServeMux()
	controller.RegisterRestController(mux)
	service.GenerateMovieList()
	http.ListenAndServe(":8080", mux)
}
