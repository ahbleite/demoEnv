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
    fmt.Fprintf(w, "Ambiente: %v\n", os.Getenv("NAMESPACE"))
}

// main function to boot up everything
func main() {
    router := mux.NewRouter()
    router.HandleFunc("/", home).Methods("GET")
    log.Fatal(http.ListenAndServe(":9999", router))
} 