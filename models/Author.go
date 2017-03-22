package models

import (
    "time"
)

type Author struct {
    Id         int        `schema:"int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID'" index:"PRIMARY KEY"`
    Name       string     `schema:"varchar(255) NOT NULL DEFAULT '' COMMENT '名称'" index:"KEY"`
    Created_at time.Time  `schema:"timestamp NULL DEFAULT NULL COMMENT '创建时间'"`
    Updated_at time.Time  `schema:"timestamp NULL DEFAULT NULL COMMENT '更新时间'"`
}
