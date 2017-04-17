package main

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
	"fmt"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello!!!")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", helloHandler).Methods("GET")

	n := negroni.Classic()
	n.Run(":8009")
}