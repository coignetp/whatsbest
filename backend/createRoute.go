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
    var resizedChoices []string
    var typeChoices []int


    for i := 0; i < requestTournament.Size; i++ {
      c, t, err := checkAndProcessChoice(requestTournament.Choices[i])
      if err == false {
        resizedChoices = append(resizedChoices, c)
        typeChoices = append(typeChoices, t)
      }
    }
    
    id = addTournament(db, requestTournament.Question, len(resizedChoices), resizedChoices, typeChoices)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("{\"id\": \"" + strconv.FormatInt(int64(id), 16) + "\"}"))
  }
}


func checkAndProcessChoice(b64 string) (string, int, bool) {
  if !strings.HasPrefix(b64, "data:image/") {
    // Raw choice
    return b64, 0, false
  }
  log.Printf("%+v is", b64)
  
  splitted := strings.SplitN(b64, ",", 2)

  // Wrong format
  if len(splitted) != 2 {
    log.Print("Wrong formatted file")
    return "", 1, true
  }

  if (strings.Contains(splitted[0], "image/jpg") || strings.Contains(splitted[0], "image/jpeg") || strings.Contains(splitted[0], "image/png") || strings.Contains(splitted[0], "image/gif")) {
    // TODO: resize ?

    return b64, 1, false
  }

  return "", 1, true
}
