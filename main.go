package main

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
	"fmt"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	pp, _ := os.Getwd()
	fmt.Fprintln(w, pp)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", helloHandler).Methods("GET")

	n := negroni.Classic()
	n.Run(":8009")
}