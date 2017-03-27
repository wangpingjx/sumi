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

    // QueryBuilder
    // DB.Table("book").Select("id, name").Where("id > ?", 0).Find()

    // Create
    // author := &Author{Name: "测试9", Created_at: time.Now(), Updated_at: time.Now()}
    // DB.Create(author)

    // Update
    // DB.Model(Author{}).Where("id = ?", 5).Update("name", "测试5")

    // Save -> update
    // var author Author
    // DB.First(&author)
    // author.Name = "测试"
    // DB.Save(&author)

    // Save -> create
    // author := Author{Name: "测试10", Created_at: time.Now(), Updated_at: time.Now()}
    // DB.Save(&author)

    // var authors []Author
    // DB.Find(&authors)
    // log.Printf("authors is: %v", authors)
    //
    // author1 := Author{}
    // DB.First(&author1)
    // log.Printf("author1 is: %v", author1)
    //
    // author2 := Author{}
    // DB.Last(&author2)
    // log.Printf("author2 is: %v", author2)

    // // 传入Slice
    // log.Println("=> for DB.First(&[]Author).......")
    // var authors2 []Author
    // DB.Find(&authors2)
    // log.Printf("authors is: %v", authors2)
    // for _, author := range authors2 {
    //     log.Printf("author is: %v", author)
    //     log.Printf("author name is: %s", author.Name)
    // }
    //
    // // 传入Slice指针
    // log.Println("=> for DB.First(&[]*Author).......")
    // var authors []*Author
    // DB.Find(&authors)
    // log.Printf("authors is: %v", authors)
    // for _, author := range authors {
    //     log.Printf("author is: %v", *author)
    //     log.Printf("author name is: %s", (*author).Name)
    // }

    // author := Author{}
    // DB.Where("id = ?", 1).Find(&author)
    // log.Printf("author is : %v", author)
    // DB.Where("id > ?", 1).First(&author)
    // log.Printf("author is : %v", author)
    // DB.Where("id > ?", 1).Last(&author)
    // log.Printf("author is : %v", author)

    // var authors []Author
    // DB.Where("id > ?", 1).Order("id DESC").Limit(2).Find(&authors)
    // log.Printf("authors is : %v", authors)
}
