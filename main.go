package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/", homeHandler).Methods(http.MethodGet)
    r.HandleFunc("/hello", helloHandler).Methods(http.MethodGet)
    r.HandleFunc("/add", addHandler).Methods(http.MethodGet)

    fmt.Println("Server started at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello World")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    fmt.Fprintf(w, "hello %s", name)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
    param1 := r.URL.Query().Get("param1")
    param2 := r.URL.Query().Get("param2")

    a := 0
    b := 0

    if _, err := fmt.Sscanf(param1, "%d", &a); err != nil {
        http.Error(w, "param1 must be an integer", http.StatusBadRequest)
        return
    }

    if _, err := fmt.Sscanf(param2, "%d", &b); err != nil {
        http.Error(w, "param2 must be an integer", http.StatusBadRequest)
        return
    }

    fmt.Fprintf(w, "%d", a+b)
}
