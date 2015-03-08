package main

import (
    "log"
    "net/http"
)

// Entry point
func main() {
	// initialize router
    router := NewRouter()
    // start server on 8080 port
    log.Fatal(http.ListenAndServe(":8080", router))
}