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

    log.Println("=> for DB.First(&Author{}).......")
    author1 := Author{}
    DB.First(&author1)
    log.Printf("author1 is: %v", author1)
    log.Printf("author1 name is: %s", author1.Name)

    author2 := Author{}
    DB.Last(&author2)
    log.Printf("author2 is: %v", author2)
    log.Printf("author2 name is: %s", author2.Name)

    // 传入Slice
    log.Println("=> for DB.First(&[]Author).......")
    var authors2 []Author
    DB.Find(&authors2)
    log.Printf("authors is: %v", authors2)
    for _, author := range authors2 {
        log.Printf("author is: %v", author)
        log.Printf("author name is: %s", author.Name)
    }

    // 传入Slice指针
    log.Println("=> for DB.First(&[]*Author).......")
    var authors []*Author
    DB.Find(&authors)
    log.Printf("authors is: %v", authors)
    for _, author := range authors {
        log.Printf("author is: %v", *author)
        log.Printf("author name is: %s", (*author).Name)
    }


}
