package common

import (
	"fmt"
	"github.com/AgentGuo/ginessential/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (error) {
	username := "root"
	password := "qwer"
	ip := "localhost"
	port := "3306"
	dbname := "ginessential"
	charset := "utf8"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True",
		username, password, ip, port, dbname, charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		return err
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil{
		return err
	}
	DB = db
	return nil
}