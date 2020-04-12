package main

import (
  "math"
  "database/sql"
)

func expectedScore(elo1 int, elo2 int) float64 {
  // Expected result of 1 against 2
  return 1.0 / (1 + math.Pow(10.0, float64(elo2 - elo1) / 400.0))
}

func updateElo(db *sql.DB, choices *ResponseChoice) {
  k := 32.0
  choices.C1.Elo += int(math.Round(k * (1.0 - float64(choices.Winner) - expectedScore(choices.C1.Elo, choices.C2.Elo))))
  choices.C2.Elo += int(math.Round(k * (float64(choices.Winner) - expectedScore(choices.C2.Elo, choices.C1.Elo))))

  setElo(db, choices.C1.Id, choices.C1.Elo)
  setElo(db, choices.C2.Id, choices.C2.Elo)
}