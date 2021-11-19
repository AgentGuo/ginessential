package controller

import (
	"github.com/AgentGuo/ginessential/common"
	"github.com/AgentGuo/ginessential/model"
	"github.com/AgentGuo/ginessential/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(c *gin.Context) {
	// get data
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	// verify data
	if len(name) == 0{
		name = utils.RandomString(10)
	}

	if len(telephone) != 11{
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg": "Your phone number must be 11 digits",
		})
		return
	}

	if isTelephoneExist(telephone){
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

	// encrypt password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg": "enc password error",
		})
		return
	}
	// create user
	common.DB.Create(&model.User{
		Name: name,
		Telephone: telephone,
		Password: string(hashedPassword),
	})

	// return result
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg": "OK",
	})
}

func Login(c *gin.Context)  {
	// get data
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	// verify data
	if len(telephone) != 11{
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg": "Your phone number must be 11 digits",
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
	var user model.User
	common.DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0{
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg": "Your phone number is not registered",
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg": "Your password error",
		})
		return
	}
	// grant token
	token := "11"

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"token":token},
		"msg": "Login succeeded",
	})
}

func isTelephoneExist(telephone string) bool{
	var user model.User
	common.DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0{
		return false
	}
	return true
}