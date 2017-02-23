package models

import (
    "goblin/goblinDB"
    "log"
)

type Book struct {
    QueryBuilder  goblinDB.QueryBuilder  // 查询构造器
}

func (this *Book) Sample() string {
    log.Println("=> in BookModel#Sample")

    sql := this.QueryBuilder.Table("book").Select("id").Where("name = ?", "test").Limit(1).ToString()
    return sql
}
