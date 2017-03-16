package models

import (
    "goblin/db"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

var DB *db.DB
var err error

func init() {
    DB, err = db.Open("mysql", "root:@tcp(127.0.0.1:3306)/sumi")
    if err != nil {
        log.Fatal(err)
    }
    DB.Migrate(&Test{})
}
