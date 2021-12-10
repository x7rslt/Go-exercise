package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"strings"
)
import "gorm.io/driver/mysql"

type User struct {
	Host        string
	Name string
	Age int
}

func main(){
	dsn := "root:admin.@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"
	db,err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "",   // 表名前缀，`User`表为`t_users`
			SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
			NameReplacer: strings.NewReplacer("CID", "Cid"), // 在转为数据库名称之前，使用NameReplacer更改结构/字段名称。
		},
	})
	if err != nil {
		log.Fatal("Mysql connect error:", err)

	}
	var user User
	db.First(&user)
	user.Name = "jinzhu 2"
	user.Age = 100
	db.Save(&user)
	fmt.Println(user)
}
