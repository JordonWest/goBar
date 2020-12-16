package data 

import (
  "database/sql"
//  "fmt"
//  "strconv"
  _ "github.com/mattn/go-sqlite3"
)

func Init() {
  db, _ :=sql.Open("sqlite3", "./booze.db")
  //checkErr(err)

  statement, _ := db.Prepare("CREATE TABEL IF NOT EXISTS beers (id INTEGER PRIMARY KEY, name TEXT, stock INTEGER)")
  statement.Exec()
}

