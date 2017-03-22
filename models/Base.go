package models

import (
    "goblin/db"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "time"
)

var DB *db.DB
var err error

func init() {
    DB, err = db.Open("mysql", "root:@tcp(127.0.0.1:3306)/sumi")
    if err != nil {
        log.Fatal(err)
    }
    // Migration
    // DB.Migrate(&Book{}, &Author{})
    // DB.DropTable(&Test{})
    // DB.ModifyColumn(&Test{}, "authorid", "int(11) unsigned  COMMENT '修改后作者ID'")
    // DB.DropColumn(&Test{}, "test1")
    // DB.AddIndex(&Test{}, "idx_authorid_created_at", "authorid", "created_at")
    // DB.AddUniqueIndex(&Test{}, "idx_name", "name")
    // DB.RemoveIndex(&Test{}, "name")

    // CRUD
    author = &Author{Name: "测试4", Created_at: time.Now(), Updated_at: time.Now()}
    DB.Create(author)
}
