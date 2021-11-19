package controller

import (
	"github.com/AgentGuo/ginessential/common"
	"github.com/AgentGuo/ginessential/model"
	"github.com/AgentGuo/ginessential/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
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

	// create user
	common.DB.Create(&model.User{
		Name: name,
		Telephone: telephone,
		Password: password,
	})

	// return result
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg": "OK",
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