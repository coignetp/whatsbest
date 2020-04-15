package main

import (
  "database/sql"
  "encoding/json"
  "log"
  "net/http"
  "strconv"

  // "mime"
  // "path/filepath"
  // "os"
)

type ResponseChoice struct {
  C1	      Choice  `json:"c1"`
  C2        Choice  `json:"c2"`
  Winner    int     `json:"winner"`
  Question  string  `json:"question"`
}

func choiceRoute(w http.ResponseWriter, r *http.Request, db *sql.DB) {
  enableCors(&w)
  if r.Method == "GET" {
    createChoice(w, r, db)
  } else if r.Method == "POST" {
    validateChoice(w, r, db)
  }
}

func NewResponseChoice(db *sql.DB, idTournament int) *ResponseChoice {
  c1, c2, found := getTwoChoices(db, idTournament)
  question := getQuestion(db, idTournament)

  if found == false {
    log.Printf("Tournament not found: %d", idTournament)
    return nil
  }

  retChoices := &ResponseChoice{
    C1: c1,
    C2: c2,
    Winner: 0,
    Question: question,
  }
  log.Printf("Selected %d and %d (%d %d)", c1.Id, c2.Id, c1.IdTournament, c2.IdTournament)

  return retChoices
}

func createChoice(w http.ResponseWriter, r *http.Request, db *sql.DB) {
  idTournament64, _ := strconv.ParseInt(r.URL.Query().Get("id"), 16, 32)
  idTournament := int(idTournament64)
  log.Print(idTournament)
  
  retChoices := NewResponseChoice(db, idTournament)
  if retChoices == nil {
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte("Not found"))
    return
  }

  jsonString, err := json.Marshal(*retChoices)
  if err != nil {
    log.Print(err)
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(jsonString)
}

func validateChoice(w http.ResponseWriter, r *http.Request, db *sql.DB) {
  decoder := json.NewDecoder(r.Body)
  var resultChoices ResponseChoice
  err := decoder.Decode(&resultChoices)

  if err != nil {
    log.Print(err.Error())
  }
  if resultChoices.Winner != 0 && resultChoices.Winner != 1 {
    log.Print("Bad winner: " + strconv.Itoa(resultChoices.Winner))
    return
  }

  // TODO: Update local elo from DB before this update
  updateElo(db, &resultChoices)

  log.Printf("New elos: %d | %d", resultChoices.C1.Elo, resultChoices.C2.Elo)
  idTournament := resultChoices.C1.IdTournament
  retChoices := NewResponseChoice(db, idTournament)
  if retChoices == nil {
    w.Write([]byte("Not found"))
    return
  }

  jsonString, err := json.Marshal(*retChoices)
  if err != nil {
    log.Print(err)
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(jsonString)
}