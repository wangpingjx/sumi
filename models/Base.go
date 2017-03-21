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

    // DB.Migrate(&Test{})

    // DB.DropTable(&Test{})

    // DB.ModifyColumn(&Test{}, "authorid", "int(11) unsigned  COMMENT '修改后作者ID'")

    // DB.DropColumn(&Test{}, "test1")

    // DB.AddIndex(&Test{}, "idx_authorid_created_at", "authorid", "created_at")

    // DB.AddUniqueIndex(&Test{}, "idx_name", "name")

    // DB.RemoveIndex(&Test{}, "name")
}
