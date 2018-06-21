package main

import (
	"net/http"
	"fmt"
	"log"
	"io"

	"github.com/gorilla/mux"

)

func main() {
	router := mux.NewRouter()

	router.Use(loggingMiddleware)
	//router.Host("www.example.com")
	router.Methods("GET", "POST")

	router.HandleFunc("/", Index)
	router.HandleFunc("/health", HealthCheckHandler)
	router.HandleFunc("/hello/{name}", Hello)

	log.Fatal(http.ListenAndServe(":3000", router))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"alive": true}`)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
	fmt.Println(r)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello %v\n", vars["name"])
}