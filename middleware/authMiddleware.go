package middleware

import (
	"github.com/AgentGuo/ginessential/common"
	"github.com/AgentGuo/ginessential/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get authorization header
		tokenString := ctx.GetHeader("Authorization")

		// validate token format
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer"){
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg": "unauthorized",
			})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid{
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg": "unauthorized",
			})
			ctx.Abort()
			return
		}
		 // get user
		 user := model.User{}
		 common.DB.First(&user, claims.UserID)
		 if user.ID == 0{
			 ctx.JSON(http.StatusUnauthorized, gin.H{
				 "code": 401,
				 "msg": "unauthorized",
			 })
			 ctx.Abort()
			 return
		 }
		 ctx.Set("user", user)
		 ctx.Next()
	}
}