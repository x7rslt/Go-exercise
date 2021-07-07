package model

import (
	"fmt"
	"food/MyLog"
	"github.com/jinzhu/gorm"
	// MySQL driver.
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

type DataBase struct {
	MyDB *gorm.DB
}

var DB *DataBase

func (db *DataBase) Init() {
	DB = &DataBase{
		MyDB: GetMySqlDB(),
	}
}

func (db *DataBase) Close() {
	DB.MyDB.Close()
}
func GetMySqlDB() *gorm.DB {
	return InitSelfDB()
}
func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"Local",
	)
	fmt.Println("Model Mysql Config:",config)
	db, err := gorm.Open("mysql", config)
	if err != nil {
		MyLog.Log.Errorf("DataBase connnect failed.DataBase name:%s,Error:%s", name, err.Error())
	}
	setupDB(db)
	db.SingularTable(true)
	return db

}

func setupDB(db *gorm.DB) {
	db.DB().SetMaxIdleConns(5)

}

func InitSelfDB() *gorm.DB {
	db := openDB(viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.name"),
	)
	return db
}
