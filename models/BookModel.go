package models

import (
    "goblin/db"
    // "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

type Book struct {
    // QueryBuilder  db.QueryBuilder  // 查询构造器
    id       int
    username string
    password string
}

func (this *Book) Sample() string {
    log.Println("=> in BookModel#Sample")

    db, err := db.Open("mysql", "root:@tcp(127.0.0.1:3306)/sumi")
    if err != nil {
        log.Fatal(err)
    }
    rows,err := db.DB().Query("select id, username, password from users where id = ?", 1)
    if err != nil {
        log.Fatal(err)
    }
    var (
        id int
        username string
        password string
    )
    for rows.Next() {
    	err := rows.Scan(&id, &username, &password)
    	if err != nil {
    		log.Fatal(err)
    	}
        log.Printf("id: %d", id)
    	log.Printf("useranme: %s", username)
        log.Printf("password: %s", password)
    }
    defer rows.Close()

    // sql := this.QueryBuilder.Table("book").Select("id").Where("name = ?", "test").Limit(1).ToString()
    return "success"
}
