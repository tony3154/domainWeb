package models

import (
	_ "github.com/go-sql-driver/mysql"
)

// 用户信息
type Users struct {
	Id   int
	Name string `orm:"unique"` // 用户名唯一
	Pwd  string
}
