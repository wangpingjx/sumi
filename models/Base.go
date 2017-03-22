package models

import (
    "goblin/db"
    _ "github.com/go-sql-driver/mysql"
    "log"
    // "time"
)

var DB *db.DB
var err error

func init() {
    DB, err = db.Open("mysql", "root:@tcp(127.0.0.1:3306)/sumi?parseTime=true")
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
    // author := &Author{Name: "测试4", Created_at: time.Now(), Updated_at: time.Now()}
    // DB.Create(author)

    author    := Author{}
    rows, err := DB.First(&author)

    // 希望达成的效果
    // log.Printf("author name is: %v", author.Name)

    // 现在的效果
    if err != nil {
        log.Fatal(err)
    }
    for rows.Next() {
        err := rows.Scan(&author.Id, &author.Name, &author.Created_at, &author.Updated_at)
        if err != nil {
            log.Fatal(err)
        }
    }
    log.Printf("result: %v", author)
}
