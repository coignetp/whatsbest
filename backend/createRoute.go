package main

import (
  "database/sql"
  "encoding/json"
  "log"
  "net/http"
  "strconv"
  "strings"
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
    var id int

    if requestTournament.Type == 1 {
      var resizedChoices []string

      for i := 0; i < requestTournament.Size; i++ {
        resizedChoices = append(resizedChoices, checkAndResizeImage(requestTournament.Choices[i]))
      }
      id = addTournament(db, requestTournament.Question, requestTournament.Size, resizedChoices, requestTournament.Type)
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("{\"id\": \"" + strconv.FormatInt(int64(id), 16) + "\"}"))
  }
}


func checkAndResizeImage(b64 string) string {
  splitted := strings.SplitN(b64, ",", 2)

  // Wrong format
  if len(splitted) != 2 {
    log.Print("Wrong formatted file")
    return ""
  }

  if (strings.Contains(splitted[0], "image/jpg") || strings.Contains(splitted[0], "image/jpeg") || strings.Contains(splitted[0], "image/png") || strings.Contains(splitted[0], "image/gif")) {
    // TODO: resize

    return b64
  }

  return ""
}
