package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func routeHandler() {
	router := mux.NewRouter().StrictSlash(true)

	// Create a route handler
	router.HandleFunc("/", HomeHandler)
	log.Fatal(http.ListenAndServe(":3000", router))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("UP!"))
}

func main() {
	fmt.Println("Starting Tweeter API on port 3000...")
	routeHandler()
}
