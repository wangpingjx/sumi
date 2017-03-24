package models

import (
    // "log"
    "time"
)

type Book struct {
    id         int        `schema:"int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID'" index:"PRIMARY KEY"`
    name       string     `schema:"varchar(255) NOT NULL DEFAULT '' COMMENT '书名'" index:"KEY"`
    desc       string     `schema:"text COMMENT '描述'"`
    author_id  int        `schema:"int(11) NOT NULL DEFAULT 0 COMMENT '作者ID'" index:"KEY"`
    created_at time.Time  `schema:"timestamp NULL DEFAULT NULL COMMENT '创建时间'"`
    updated_at time.Time  `schema:"timestamp NULL DEFAULT NULL COMMENT '更新时间'"`
}

// func (this *Book) List(perpage string, page string) []map[string]interface{} {
func (this *Book) List(perpage string, page string) string {

    return "list"
    // rows, err := DB.Table("book").Select("id, name").Where("id > ?", 0).Find()
    // if err != nil {
    //     log.Fatal(err)
    // }
    // var (
    //     id   int
    //     name string
    // )
    // var data []map[string]interface{}
    // for rows.Next() {
    //     err := rows.Scan(&id, &name)
    //     if err != nil {
    //         log.Fatal(err)
    //     }
    //     tmp := map[string]interface{}{"id": id, "name": name}
    //     data = append(data, tmp)
    // }
    // return data
}


// func (this *Book) Show(bookId string) map[string]interface{} {
func (this *Book) Show(bookId string) string {
    return "show"
    // rows, err := DB.Table("book").Select("id, name").Where("id = ?", bookId).Find()
    // if err != nil {
    //     log.Fatal(err)
    // }
    // data := make(map[string]interface{})
    // var (
    //     id   int
    //     name string
    // )
    // for rows.Next() {
    // 	err := rows.Scan(&id, &name)
    // 	if err != nil {
    // 		log.Fatal(err)
    // 	}
    //     data["id"]   = id
    //     data["name"] = name
    //     break
    // }
    // return data
}
