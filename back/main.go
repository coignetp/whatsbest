package main

import (
  // "database/sql"
  "fmt"
  // "html"
  "log"
  "net/http"
  "math/rand"
  "time"
)


func createTournament() int64 {
  var id = rand.Int63()

  log.Print(id)

  return id
}


func main() {
  log.Print("Listening :8081")
  rand.Seed(time.Now().UTC().UnixNano())

  http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
    // fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    log.Print("Creating a new tournament")
    
    fmt.Fprintf(w, "%d", createTournament())
  })

  log.Fatal(http.ListenAndServe(":8081", nil))
}