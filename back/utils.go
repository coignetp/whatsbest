package main

import (
	"database/sql"
	"net/http"
)

func enableCors(w *http.ResponseWriter) {
  // TODO: restrict
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
  (*w).Header().Set("Access-Control-Allow-Headers", "*")
  (*w).Header().Set("Access-Control-Allow-Method", "POST,GET")
}


func makeDbClientHandler(fn func(http.ResponseWriter, *http.Request, *sql.DB), c *sql.DB) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    fn(w, r, c)
  }
}