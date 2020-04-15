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
  port := "8081"
  if len(os.Args) >= 2 {
    port = os.Args[1]
  }
  log.Printf("Listening :%s", port)
  rand.Seed(time.Now().UTC().UnixNano())

  fs := http.FileServer(http.Dir("./web/dist"))
  http.Handle("/", fs)
  
  // Clear database
  os.Remove("tournament.db")
  _, err := os.Stat("tournament.db")

  if os.IsNotExist(err) {
    file, err := os.Create("tournament.db")
    if err != nil {
      log.Fatal(err.Error())
    }
    file.Close()
  }

  sqliteDatabase, err := sql.Open("sqlite3", "./tournament.db")
  if err != nil {
    log.Fatal(err.Error())
  }
  defer sqliteDatabase.Close()

  createTables(sqliteDatabase)

  http.HandleFunc("/api/create", makeDbClientHandler(createTournamentRoute, sqliteDatabase))
  http.HandleFunc("/api/choice", makeDbClientHandler(choiceRoute, sqliteDatabase))
  http.HandleFunc("/api/result", makeDbClientHandler(resultRoute, sqliteDatabase))

  log.Fatal(http.ListenAndServe(":" + port, nil))
}