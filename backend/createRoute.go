package main

import (
  "database/sql"
  "encoding/json"
  "log"
  "net/http"
  "strconv"
)

type RequestTournament struct {
  Question  string    `json:"question"`
  Size      int       `json:"size"`
  Type      int       `json:"type"`
  Choices   []string  `json:"choices"`
}

func createTournamentRoute(w http.ResponseWriter, r *http.Request, db *sql.DB) {
  enableCors(&w)
  
  if r.Method == "POST" {
    decoder := json.NewDecoder(r.Body)
    var requestTournament RequestTournament
    err := decoder.Decode(&requestTournament)

    if err != nil {
      log.Print(err.Error())
    }

    log.Println("Question: " + requestTournament.Question) 

    id := addTournament(db, requestTournament.Question, requestTournament.Size, requestTournament.Choices, requestTournament.Type)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("{\"id\": \"" + strconv.FormatInt(int64(id), 16) + "\"}"))
  }
}
