package db

import (
	"database/sql"
	"strings"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
)

var DB *sql.DB

var (
	username = "root"
	password = "zy123456lk"
	ip = "127.0.0.1"
	port = "3306"
	dbname = "user"
)

func init() {
	path := strings.Join([]string{username, ":", password, "@tcp(", ip, ":", port, ")/", dbname, "?charset=utf8"}, "")
	DB, _ = sql.Open("mysql", path)
	if err := DB.Ping(); err != nil {
		fmt.Println("数据库连接失败", err)
	}
}



