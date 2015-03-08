package main

import "time"

// Model for one todo item
type Todo struct {
    Name      string    `json:"name"`
    Completed bool      `json:"completed"`
    Due       time.Time `json:"due"`
}

// Model for todo's collection
type Todos []Todo
