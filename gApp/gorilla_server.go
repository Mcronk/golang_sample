package main

import (
//	"fmt"
    "net/http"
    "log"
    "github.com/gorilla/mux"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
    writer.Write([]byte("The server is working\n"))
}

func IndexHandler(writer http.ResponseWriter, request *http.Request) {
    writer.Write([]byte("Index Page\n"))
}

func main() {
    request := mux.NewRouter()
    request.HandleFunc("/", Handler)
    request.HandleFunc("/index", IndexHandler )

    log.Fatal(http.ListenAndServe(":8080", request))
}