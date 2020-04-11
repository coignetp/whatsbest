package main

import (
	"database/sql"
	"log"
	"math/rand"
)

type Tournament struct {
	id				int64
	question	string
	size			int
	choices		[][]byte
	elos			[]int
}

func createTables(db *sql.DB) {
	createTournamentSQL := `CREATE TABLE IF NOT EXISTS tournaments (
    "id" INTEGER PRIMARY KEY,
    "question" TEXT,
    "size" INTEGER
  );`
  createChoiceSQL := `CREATE TABLE IF NOT EXISTS choices (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "type" INTEGER,
    "bytestream" TEXT,
    "elo" INTEGER
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

	log.Println("Creating Choices table")
	statement, err = db.Prepare(createChoiceSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()

	log.Println("Creating TournamentChoice table")
	statement, err = db.Prepare(createTournamentToChoiceSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
}

func addTournament(db *sql.DB, question string, size int, choices [][]byte, choicesType int) int64 {
	var id = rand.Int63()
	baseElo := 1000
	// TODO: check id is available
	// TODO: remove
	log.Print(id)

	if len(choices) != size {
		return -1
	}

	tx, _ := db.Begin()
	statement, _ := tx.Prepare("INSERT INTO tournaments (id, question, size) values (?,?,?)")
	_, err := statement.Exec(id, question, size)
	if err != nil {
		log.Print(err.Error())
	}

	for i := 0; i < size; i++ {
		choiceStatement, _ := tx.Prepare("INSERT INTO choices (type, bytestream, elo) values (?,?,?)")
		res, err := choiceStatement.Exec(choicesType, choices[i], baseElo)
		if err != nil {
			log.Print(err.Error())
		}

		choiceId, _ := res.LastInsertId()

		tournamentChoiceStatement, _ := tx.Prepare("INSERT INTO tournamentchoice (idTournament, idChoice, choiceNumber) values (?,?,?);")
		_, err = tournamentChoiceStatement.Exec(id, choiceId, i)
		if err != nil {
			log.Print(err.Error())
		}
	}

	tx.Commit()

	return id
}