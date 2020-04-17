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

  _ "github.com/lib/pq"
)

func main() {
  port := flag.Int("port", 8081, "The exposed port")
  devMode := flag.Bool("dev", false, "If the server should be started in development mode")
  flag.Parse()

  log.Printf("Listening :%d", *port)
  log.Printf("Dev mode is %+v", *devMode)
  rand.Seed(time.Now().UTC().UnixNano())

  // In porduction (on heroku), the golang backend serves the frontend
  // due to the free-tier limitation
  if !*devMode {
    fs := http.FileServer(http.Dir("./web/dist"))
    http.Handle("/", fs)
  }

  database, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
  if err != nil {
    log.Fatal(err.Error())
  }
  defer database.Close()

  createTables(database)

  // Create the different routes
  http.HandleFunc("/api/create", makeDbClientHandler(createTournamentRoute, database))
  http.HandleFunc("/api/choice", makeDbClientHandler(choiceRoute, database))
  http.HandleFunc("/api/result", makeDbClientHandler(resultRoute, database))

  log.Fatal(http.ListenAndServe(":" + strconv.Itoa(*port), nil))
}