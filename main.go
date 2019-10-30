package main

import (
	"log"
	"net/http"

	"example-rest-api/dummy_db"
	"example-rest-api/handlers"

	"github.com/gorilla/mux"
)

func main() {

	dummy_db.InitDB()

	r := mux.NewRouter().StrictSlash(false)

	r.HandleFunc("/{url}", handlers.RedirectToRealURL).Methods("GET")

	apiRouter := r.PathPrefix("/api/v0").Subrouter()

	apiRouter.HandleFunc("/addlink", handlers.AddLink).Methods("POST")

	apiRouter.HandleFunc("/link/{id}", handlers.GetLink).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))

}
