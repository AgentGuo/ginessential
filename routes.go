package main

import (
	"github.com/AgentGuo/ginessential/controller"
	"github.com/AgentGuo/ginessential/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine){
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
}
