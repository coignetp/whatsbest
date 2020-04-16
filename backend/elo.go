package main

import (
  "database/sql"
  "math"
  "log"
)

// Return expected result of 1 against 2
// The more the returned value is closed to 1.0, the more 
// 1 should win against 2.
func expectedScore(elo1 int, elo2 int) float64 {
  return 1.0 / (1 + math.Pow(10.0, float64(elo2 - elo1) / 400.0))
}

// Update elos in the database according to a winner of a match
func updateElo(db *sql.DB, choices *ResponseChoice) {
  k := 32.0
  // Get the most recent elos first
  choices.C1.Elo = mostRecentElo(db, choices.C1.Id)
  choices.C2.Elo = mostRecentElo(db, choices.C2.Id)

  if (choices.C1.Elo == -1 || choices.C2.Elo == -1) {
    return
  }

  gap := int(math.Round(k * (1.0 - float64(choices.Winner) - expectedScore(choices.C1.Elo, choices.C2.Elo))))
  choices.C1.Elo += gap
  choices.C2.Elo -= gap

  log.Printf("Elo changement: %d", gap)

  setElo(db, choices.C1.Id, choices.C1.Elo)
  setElo(db, choices.C2.Id, choices.C2.Elo)
}