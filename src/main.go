package main

import (
	controller "crud/src/pkg/controller"
	"net/http"
)

func main() {
	var mux *http.ServeMux = http.NewServeMux()
	controller.RegisterRestController(mux)
	http.ListenAndServe(":8080", mux)

}
