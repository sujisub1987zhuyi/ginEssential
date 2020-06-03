package common

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"jim.zhu/ginessential/model"
)

var DB *gorm.DB

//InitDB init db
func InitDB() *gorm.DB {
	dirvername := "mysql"
	host := "localhost"
	port := "3306"
	database := "ginessential"
	username := "root"
	password := "root"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(dirvername, args)
	if err != nil {
		panic("fail to connect database , err: " + err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}
