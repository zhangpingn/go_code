package db

import (
	"fmt"
	"go_code/execrise/xmlexec/xmlexec03/parsexml"
	"database/sql"
	"strings"
	"go_code/execrise/xmlexec/xmlexec03/model"
	_ "github.com/Go-SQL-Driver/MySQL"
)

const (
    userName = "root"
    password = "zy123456lk"
    ip = "127.0.0.1"
    port = "3306"
    dbName = "student"
)

var DB *sql.DB

func GetData(file string, sqlId string, sid string) interface{} {
	var student model.Student
	obj := parsexml.Parse(file, sqlId)
	if obj == nil {
		fmt.Println("获取sql语句失败")
		return nil
	}
	path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	DB, _ = sql.Open("mysql", path)

	if err := DB.Ping(); err != nil {
		fmt.Println("数据库连接失败", err)
		return nil
	}
	//设置mysql数据库最大连接数
	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(20)
	
	selectobj, ok :=obj.(parsexml.Select)
	if ok {
		// tx, err := DB.Begin()
		// if err != nil {
		// 	fmt.Println("开启事务失败", err)
		// }
		// stmt, err := tx.Prepare(selectobj.Comment)
		// if err != nil {
		// 	fmt.Println("sql执行异常", err)
		// 	return nil
		// }
		// res, err := stmt.Exec(sid)
		// if err != nil {
		// 	fmt.Println("sql执行异常", err)
		// 	return nil
		// }
		// fmt.Println(res)
		// tx.Commit()
		fmt.Println(selectobj.Comment)
		err := DB.QueryRow(selectobj.Comment, sid).Scan(&student.Sid, &student.Sname, &student.Age, &student.Score)
	
		if err != nil {
			fmt.Println("查询出错了")
		}
	}
	
	return student
}