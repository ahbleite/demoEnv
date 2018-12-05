package main

import (
    "github.com/gorilla/mux"
    "log"
	"net/http"
	"fmt"
	"os"
)

func home(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Ambiente: %v\nThundercats, HOOOOOOO!!!!", os.Getenv("ENVIRONMENT"))
}

func health(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "")
}

// main function to boot up everything
func main() {
    router := mux.NewRouter()
    router.HandleFunc("/", home).Methods("GET")
    router.HandleFunc("/health", health).Methods("GET")
    log.Fatal(http.ListenAndServe(":9999", router))
} 