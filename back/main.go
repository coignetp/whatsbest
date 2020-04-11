package main

import (
  "database/sql"
  "log"
  "net/http"
  "math/rand"
  "time"
  "os"

  _ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

func main() {
  log.Print("Listening :8081")
  rand.Seed(time.Now().UTC().UnixNano())
  
  // Clear database
  os.Remove("tournament.db")
  file, err := os.Create("tournament.db")
  if err != nil {
    log.Fatal(err.Error())
  }
  file.Close()

  sqliteDatabase, err := sql.Open("sqlite3", "./tournament.db")
  if err != nil {
    log.Fatal(err.Error())
  }
  defer sqliteDatabase.Close()

  createTables(sqliteDatabase)

  http.HandleFunc("/create", makeDbClientHandler(createTournament, sqliteDatabase))

  http.HandleFunc("/redirect", func (w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "http://localhost:8080/mash/3", 303)
  })

  log.Fatal(http.ListenAndServe(":8081", nil))
}