package models

import(
    // "log"
    "time"
)

type User struct {
    Id         int        `schema:"int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID'" index:"PRIMARY KEY"`
    Name       string     `schema:"varchar(255) NOT NULL DEFAULT '' COMMENT '用户名'" index:"UNIQUE KEY"`
    Password   string     `schema:"varchar(20) NOT NULL DEFAULT '' COMMENT '密码'"`
    Created_at time.Time  `schema:"timestamp NULL DEFAULT NULL COMMENT '创建时间'"`
    Updated_at time.Time  `schema:"timestamp NULL DEFAULT NULL COMMENT '更新时间'"`
}
