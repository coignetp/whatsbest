package main

import (
  "database/sql"
  "encoding/json"
  "log"
  "net/http"
  "strconv"
)

func resultRoute(w http.ResponseWriter, r *http.Request, db *sql.DB) {
  enableCors(&w)

  if r.Method == "GET" {
    idTournament64, _ := strconv.ParseInt(r.URL.Query().Get("id"), 16, 32)
    idTournament := int(idTournament64)
    log.Print(idTournament)

    result := getAllChoices(db, idTournament)
    question := getQuestion(db, idTournament)

    if len(result) == 0 {
      w.WriteHeader(http.StatusNotFound)
      w.Write([]byte("Not found"))
      return
    }

    jsonResult, err := json.Marshal(result)
    if err != nil {
      log.Print(err)
    }

    jsonString := []byte(
      `{"question": "` + question + `", "result": ` + string(jsonResult) + "}")

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonString)
  }
}