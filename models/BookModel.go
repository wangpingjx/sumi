package models

import (
    "goblin/db"
    // "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

/* TODO 默认表名为类名的小写，也可通过 TableName()方法指定表名 */
type Book struct {
    id       int
    username string
    password string
}

/* TODO 支持设置表名 */
/*
func (this Book) TableName() string {
    return "books"
}
*/
func (this *Book) Sample() string {
    log.Println("=> in BookModel#Sample")

    db, err := db.Open("mysql", "root:@tcp(127.0.0.1:3306)/sumi")
    if err != nil {
        log.Fatal(err)
    }
    rows, err := db.Table("users").Select("id, username, password").Where("id in (?)", []int{1,2}).Where("username like ?", "%es%").Find()
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
    return "success"
}
