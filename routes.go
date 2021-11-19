package main

import (
	"github.com/AgentGuo/ginessential/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine){
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
}
