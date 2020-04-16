package main

import (
  "database/sql"
  "encoding/json"
  "log"
  "net/http"
  "strconv"
)

// The response sent to the frontend when asking for a new choice
// The same struct is used to retrieve the winner from the frontend
type ResponseChoice struct {
  C1	      Choice  `json:"c1"`
  C2        Choice  `json:"c2"`
  Winner    int     `json:"winner"`
  Question  string  `json:"question"`
}

// POST request: process a choice made in the body and send a new choice to make
// GET request: just send a choice to make
func choiceRoute(w http.ResponseWriter, r *http.Request, db *sql.DB) {
  enableCors(&w)
  if r.Method == "GET" {
    createChoice(w, r, db)
  } else if r.Method == "POST" {
    validateChoice(w, r, db)
  }
}

// Return a new choice to make between 2 possibilities
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

// Handle de /api/choice GET requests
// Just send back a new choice to make
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

// Handle de /api/choice POST requests
// Update elo with the given request body
// Send back a new choice to make
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

  // Update the elos according to the winner
  updateElo(db, &resultChoices)

  log.Printf("New elos: %d | %d", resultChoices.C1.Elo, resultChoices.C2.Elo)

  // New choice to make
  idTournament := resultChoices.C1.IdTournament
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