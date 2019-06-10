package main

import (
	"go_code/execrise/xmlexec/xmlexec03/db"
	"go_code/execrise/xmlexec/xmlexec03/model"
	"fmt"
)

func main() {
	stu := db.GetData("D:/stu.xml", "getBySid", "10001")
	student := stu.(model.Student)
	fmt.Println(student)
}