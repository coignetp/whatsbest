package main

import (
  "database/sql"
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "strconv"

  // "mime"
  // "path/filepath"
  // "os"
)

type ResponseChoice struct {
  C1	    Choice  `json:"c1"`
  C2      Choice  `json:"c2"`
  Winner  int     `json:"winner"`
}

func choiceRoute(w http.ResponseWriter, r *http.Request, db *sql.DB) {
  if r.Method == "GET" {
    createChoice(w, r, db)
  } else if r.Method == "POST" {
    validateChoice(w, r, db)
  }
}

func createChoice(w http.ResponseWriter, r *http.Request, db *sql.DB) {
  enableCors(&w)

  log.Print("In create choice")

  idTournament, _ := strconv.ParseInt(r.URL.Query().Get("id"), 16, 64)
  log.Print(idTournament)
  c1, c2, found := getTwoChoices(db, idTournament)

  if found == false {
    fmt.Fprint(w, "not found")
    return
  }

  retChoices := &ResponseChoice{
    C1: c1,
    C2: c2,
    Winner: 0,
  }
  log.Printf("Selected %d and %d", c1.Id, c2.Id)

  jsonString, err := json.Marshal(retChoices)
  if err != nil {
    log.Print(err)
  }
  // log.Print(string(jsonString))

  ///
  // var fileName string = "image"
  // detectedFileType := http.DetectContentType(c1.Bytestream)
  // switch detectedFileType {
  //   case "image/jpeg", "image/jpg":
  //   case "image/gif", "image/png":
  //     break
  //   default:
  //     log.Println("Unsupported file type")
  // }

  // fileEndings, err := mime.ExtensionsByType(detectedFileType)
  // if err != nil {
  //   log.Println("3: " + err.Error())
  // }
  // newPath := filepath.Join("uploaded", fileName+fileEndings[0])
  // newFile, err := os.Create(newPath)
  // if err != nil {
  //   log.Println("4: " + err.Error())
  // }
  // defer newFile.Close()
  // _, err = newFile.Write(c1.Bytestream)
  // if err != nil {
  //   log.Println("5: " + err.Error())
  // }


  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(jsonString)
}

func validateChoice(w http.ResponseWriter, r *http.Request, db *sql.DB) {
}