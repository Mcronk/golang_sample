package main

import (
//	"fmt"
    "net/http"
    "log"
    "github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("The server is working\n"))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Index Page\n"))
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", Handler)
    r.HandleFunc("/index", IndexHandler )

    log.Fatal(http.ListenAndServe(":8000", r))
}