package main

import (
  "database/sql"
  "flag"
  "log"
  "net/http"
  "math/rand"
  "strconv"
  "time"
  "os"

  _ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

func main() {
  port := flag.Int("port", 8081, "The exposed port")
  devMode := flag.Bool("dev", false, "If the server should be started in development mode")
  flag.Parse()

  log.Printf("Listening :%d", *port)
  log.Printf("Dev mode is %+v", *devMode)
  rand.Seed(time.Now().UTC().UnixNano())

  if !*devMode {
    fs := http.FileServer(http.Dir("./web/dist"))
    http.Handle("/", fs)
  } else {
    // Clear database
    os.Remove("tournament.db")
  }

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

  log.Fatal(http.ListenAndServe(":" + strconv.Itoa(*port), nil))
}