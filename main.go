package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"time"
)

const char string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(11);not null"`
	Password string `gorm:"size:255;not null"`
}
func main()  {
	// connect to db
	db, err := InitDB()
	if err != nil{
		panic(err)
	}
	// close db
	defer func() {
		dbClose, err := db.DB()
		if err == nil{
			dbClose.Close()
		}
	}()
	// get an engine instance
	r := gin.Default()
	r.POST("/api/auth/register", func (c *gin.Context) {
		name := c.PostForm("name")
		telephone := c.PostForm("telephone")
		password := c.PostForm("password")
		// verify data
		if len(name) == 0{
			name = RandomString(10)
		}

		if len(telephone) != 11{
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg": "Your phone number must be 11 digits",
			})
			return
		}

		if isTelephoneExist(db, telephone){
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg": "Your phone number is already in use",
			})
			return
		}

		if len(password) < 6{
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg": "Your password can't less than 6 digits",
			})
			return
		}

		// create user
		db.Create(&User{
			Name: name,
			Telephone: telephone,
			Password: password,
		})

		// return result
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "OK",
		})
	})
	r.Run()
}

func isTelephoneExist(db *gorm.DB, telephone string) bool{
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0{
		return false
	}
	return true
}

func RandomString(n int) string {
	var s bytes.Buffer
	rand.NewSource(time.Now().UnixNano())
	for i := 0; i < n; i++{
		s.WriteByte(char[rand.Intn(len(char))])
	}
	return s.String()
}

func InitDB() (*gorm.DB, error) {
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
		return nil, err
	}
	db.AutoMigrate(&User{})
	return db, nil
}