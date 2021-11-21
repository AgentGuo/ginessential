package controller

import (
	"github.com/AgentGuo/ginessential/common"
	"github.com/AgentGuo/ginessential/dto"
	"github.com/AgentGuo/ginessential/model"
	"github.com/AgentGuo/ginessential/response"
	"github.com/AgentGuo/ginessential/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Register(ctx *gin.Context) {
	// get data
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	// verify data
	if len(name) == 0{
		name = utils.RandomString(10)
	}

	if len(telephone) != 11{
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil,
			"Your phone number must be 11 digits")
		return
	}

	if isTelephoneExist(telephone){
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil,
			"Your phone number is already in use")
		return
	}

	if len(password) < 6{
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil,
			"Your password can't less than 6 digits")
		return
	}

	// encrypt password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
		response.Response(ctx, http.StatusInternalServerError, 500, nil,
			"enc password error")
		return
	}
	// create user
	common.DB.Create(&model.User{
		Name: name,
		Telephone: telephone,
		Password: string(hashedPassword),
	})

	// return result
	response.Success(ctx, nil, "OK")
}

func Login(ctx *gin.Context)  {
	// get data
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	// verify data
	if len(telephone) != 11{
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil,
			"Your phone number must be 11 digits")
		return
	}
	if len(password) < 6{
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil,
			"Your password can't less than 6 digits")
		return
	}
	var user model.User
	common.DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0{
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil,
			"Your phone number is not registered")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil{
		response.Fail(ctx, nil, "Your password error")
		return
	}
	// grant token
	token, err := common.ReleaseToken(user)
	if err != nil{
		response.Response(ctx, http.StatusInternalServerError, 500, nil,
			"grant token failed")
		log.Printf("token grant failed\n")
		return
	}
	response.Success(ctx, gin.H{"token":token}, "Login succeeded")
}

func isTelephoneExist(telephone string) bool{
	var user model.User
	common.DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0{
		return false
	}
	return true
}

func Info(ctx *gin.Context)  {
	user, _ := ctx.Get("user")
	response.Success(ctx, gin.H{"user": dto.ToUserDto(user.(model.User))},
		"Successfully obtained information")
}