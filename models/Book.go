package models

import (
    "log"
    "time"
)

/* TODO 默认表名为类名的小写，也可通过 TableName()方法指定表名 */
type Book struct {
    id       int
    username string
    password string
}

type Test struct {
    id         int       `db:"int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID'" index:"PRIMARY KEY"`
    authorId   string    `db:"int(11) unsigned  COMMENT '作者ID'" index:"INDEX"`
    name       string    `db:"varchar(255) NOT NULL DEFAULT '' COMMENT '名称'" index:"UNIQUE"`
    desc       string    `db:"text  COMMENT '描述'"`
    created_at time.Time `db:"timestamp NULL DEFAULT NULL"`
    updated_at time.Time `db:"timestamp NULL DEFAULT NULL"`
}

func (this *Book) List(perpage string, page string) []map[string]interface{} {
    rows, err := DB.Table("book").Select("id, name").Where("id > ?", 0).Find()
    if err != nil {
        log.Fatal(err)
    }
    var (
        id   int
        name string
    )
    var data []map[string]interface{}
    for rows.Next() {
        err := rows.Scan(&id, &name)
        if err != nil {
            log.Fatal(err)
        }
        tmp := map[string]interface{}{"id": id, "name": name}
        data = append(data, tmp)
    }
    return data
}


func (this *Book) Show(bookId string) map[string]interface{} {
    rows, err := DB.Table("book").Select("id, name").Where("id = ?", bookId).Find()
    if err != nil {
        log.Fatal(err)
    }
    data := make(map[string]interface{})
    var (
        id   int
        name string
    )
    for rows.Next() {
    	err := rows.Scan(&id, &name)
    	if err != nil {
    		log.Fatal(err)
    	}
        data["id"]   = id
        data["name"] = name
        break
    }
    return data
}
