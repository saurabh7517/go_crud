package main

import (
	restwrapper "crud/src/pkg/controller"
	"net/http"
)

func main() {
	var mux *http.ServeMux = http.NewServeMux()
	restwrapper.RegisterRestController(mux)
	http.ListenAndServe(":8080", mux)

}
