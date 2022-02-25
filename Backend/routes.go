package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes() *chi.Mux {
	mux := chi.NewMux()
	mux.Get("/", getAllPrograms)
	mux.Post("/program", safeProgram)
	mux.Get("/programs", getAllPrograms)
	mux.Get("/program/:id", getProgramById)
	mux.Get("/run", runProgram)

	return mux
}

func safeProgram(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("done-by", "tomas")

	res := map[string]interface{}{"message": "hello world"}

	_ = json.NewEncoder(w).Encode(res)
}

func getAllPrograms(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is my server")
}

func getProgramById(w http.ResponseWriter, r *http.Request) {

}

func runProgram(w http.ResponseWriter, r *http.Request) {

}
