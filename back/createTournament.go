package main

import (
  "database/sql"
  "fmt"
  // "html"
  "log"
  "net/http"
  "io/ioutil"
  "mime"
  "os"
  "path/filepath"
  "strconv"
)

func createTournament(w http.ResponseWriter, r *http.Request, db *sql.DB) {
  var maxUploadSize int64 = 1024 * 1024 * 16
  
  if r.Method == "POST" {
    enableCors(&w)

    // _ = r.ParseMultipartForm(maxUploadSize)
    // file, handler, err := r.FormFile("image")
    r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
    err := r.ParseMultipartForm(maxUploadSize)
    if err != nil {
      log.Println("Error retriving data")
      log.Println(err)
    }
    
    baseName := "image"
    numberOfChoices, _ := strconv.Atoi(r.FormValue("size"))
    if numberOfChoices < 2 {
      fmt.Fprintf(w, "%d", -1)
      return
    }

    choices := make([][]byte, numberOfChoices)
    for i := 0; i < numberOfChoices; i++ {
      choices[i] = retrieveImage(r, baseName+strconv.Itoa(i))
    }

    question := r.FormValue("question")
    log.Println("Question: " + question) 

    id := addTournament(db, question, numberOfChoices, choices, 1)

    fmt.Fprint(w, strconv.FormatInt(id, 16))
  }
}

func retrieveImage(r *http.Request, name string) []byte {
  log.Println("Retrieving " + name)
  file, _, err := r.FormFile(name)
  if err != nil {
    log.Print("1: " + err.Error())
  }
  defer file.Close()
  fileBytes, err := ioutil.ReadAll(file)
  if err != nil {
    log.Print("2: " + err.Error())
  }

  var fileName string = name
  detectedFileType := http.DetectContentType(fileBytes)
  switch detectedFileType {
    case "image/jpeg", "image/jpg":
    case "image/gif", "image/png":
      break
    default:
      log.Println("Unsupported file type")
  }

  fileEndings, err := mime.ExtensionsByType(detectedFileType)
  if err != nil {
    log.Println("3: " + err.Error())
  }
  newPath := filepath.Join("uploaded", fileName+fileEndings[0])
  newFile, err := os.Create(newPath)
  if err != nil {
    log.Println("4: " + err.Error())
  }
  defer newFile.Close()
  _, err = newFile.Write(fileBytes)
  if err != nil {
    log.Println("5: " + err.Error())
  }

  return fileBytes
}