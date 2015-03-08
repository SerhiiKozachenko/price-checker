package main

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

// Handlers
// GET / 
func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

// GET /todos
func TodoIndex(w http.ResponseWriter, r *http.Request) {
    todos := Todos{
        Todo{Name: "Write presentation"},
        Todo{Name: "Host meetup"},
    }

    s := "test"

    t := json.NewEncoder(w).Encode(s);

    if (t != nil) {
        panic(t)
    }

    if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }
}

// GET /todos/{todoId}
func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}
