package main

import (
  "database/sql"
  "log"
  "math/rand"
  _ "encoding/json"
  "strconv"
)

type Choice struct {
  Id            int     `json:"id"`
  IdTournament  int     `json:"idtournament"`
  Ctype         int     `json:"type"`
  Bytestream    string  `json:"bytestream"`
  Elo           int     `json:"elo"`
  Totalround    int     `json:"totalround"`
}

func createTables(db *sql.DB) {
  createTournamentSQL := `CREATE TABLE IF NOT EXISTS tournaments (
    "id" INTEGER PRIMARY KEY,
    "question" TEXT,
    "size" INTEGER
  );`
  createChoiceSQL := `CREATE TABLE IF NOT EXISTS choices (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "idTournament" INTEGER,
    "type" INTEGER,
    "bytestream" BLOB,
    "elo" INTEGER,
    "totalround" INTEGER
  );`
  createTournamentToChoiceSQL := `CREATE TABLE IF NOT EXISTS tournamentchoice(
    "idTournament" INTEGER,
    "idChoice" INTEGER,
    "choiceNumber" INTEGER,
    PRIMARY KEY (idTournament, idChoice)
	);`

  log.Println("Creating Tournaments table")
  statement, err := db.Prepare(createTournamentSQL)
  if err != nil {
    log.Fatal(err.Error())
  }
  statement.Exec()
  defer statement.Close()

  log.Println("Creating Choices table")
  statement, err = db.Prepare(createChoiceSQL)
  if err != nil {
    log.Fatal(err.Error())
  }
  statement.Exec()
  defer statement.Close()

  log.Println("Creating TournamentChoice table")
  statement, err = db.Prepare(createTournamentToChoiceSQL)
  if err != nil {
    log.Fatal(err.Error())
  }
  statement.Exec()
  defer statement.Close()
}

func addTournament(db *sql.DB, question string, size int, choices []string, choicesType int) int {
  var id = rand.Int31()
  baseElo := 1000
  baseRound := 0
  // TODO: check id is available
  // TODO: remove
  log.Print(id)

  if len(choices) != size {
    return -1
  }

  tx, _ := db.Begin()
  statement, _ := tx.Prepare("INSERT INTO tournaments (id, question, size) values (?,?,?);")
  _, err := statement.Exec(id, question, size)
  if err != nil {
    log.Print(err.Error())
  }
  defer statement.Close()

  for i := 0; i < size; i++ {
    choiceStatement, _ := tx.Prepare("INSERT INTO choices (idTournament, type, bytestream, elo, totalround) values (?,?,?,?,?);")
    res, err := choiceStatement.Exec(id, choicesType, choices[i], baseElo, baseRound)
    if err != nil {
      log.Print(err.Error())
    }
    defer choiceStatement.Close()

    choiceId, _ := res.LastInsertId()

    tournamentChoiceStatement, _ := tx.Prepare("INSERT INTO tournamentchoice (idTournament, idChoice, choiceNumber) values (?,?,?);")
    _, err = tournamentChoiceStatement.Exec(id, choiceId, i)
    if err != nil {
      log.Print(err.Error())
    }
    defer tournamentChoiceStatement.Close()
  }

  err = tx.Commit()
  if err != nil {
    log.Print(err.Error())
  }

  return int(id)
}

func getTwoChoices(db *sql.DB, idTournament int) (Choice, Choice, bool) {
  strId := strconv.Itoa(idTournament)
  query := "SELECT * FROM choices WHERE idTournament=" + strId + " ORDER BY RANDOM() LIMIT 2;"

  choices, err := db.Query(query)
  if err != nil {
    log.Print(err.Error())
  }
  defer choices.Close()

  var c1 Choice
  var c2 Choice

  if choices.Next() == false {
    return c1, c2, false
  }
  choices.Scan(&c1.Id, &c1.IdTournament, &c1.Ctype, &c1.Bytestream, &c1.Elo, &c1.Totalround)

  if choices.Next() == false {
    return c1, c2, false
  }
  choices.Scan(&c2.Id, &c2.IdTournament, &c2.Ctype, &c2.Bytestream, &c2.Elo, &c2.Totalround)

  return c1, c2, true
}

func getAllChoices(db *sql.DB, idTournament int) []Choice {
  strId := strconv.Itoa(idTournament)
  query := "SELECT * FROM choices WHERE idTournament=" + strId + " ORDER BY elo DESC;"

  choices, err := db.Query(query)
  if err != nil {
    log.Print(err.Error())
  }
  defer choices.Close()

  var result []Choice
  for choices.Next() {
    var c Choice
    choices.Scan(&c.Id, &c.IdTournament, &c.Ctype, &c.Bytestream, &c.Elo, &c.Totalround)
    result = append(result, c)
  }

  return result
}

func setElo(db *sql.DB, idChoice, newElo int) bool {
  tx, _ := db.Begin()

  statement, err := tx.Prepare("UPDATE choices SET elo=? WHERE id=?;")
  if err != nil {
    log.Print(err.Error())
  }
  _, err = statement.Exec(newElo, idChoice)
  defer statement.Close()
  if err != nil {
    log.Print(err.Error())
    return false
  }

  tx.Commit()

  return true
}

func getQuestion(db *sql.DB, idTournament int) string {
  strId := strconv.Itoa(idTournament)
  query := "SELECT question FROM tournaments WHERE id=" + strId + ";"

  tournament, err := db.Query(query)
  if err != nil {
    log.Print(err.Error())
  }
  defer tournament.Close()

  if tournament.Next() {
    var question string
    tournament.Scan(&question)

    return question
  }

  return "Unknown"
}