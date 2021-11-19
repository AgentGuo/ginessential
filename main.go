package main

import (
	"github.com/AgentGuo/ginessential/common"
	"github.com/gin-gonic/gin"
)

func main()  {
	// connect to db
	err := common.InitDB()
	if err != nil{
		panic(err)
	}
	// close db
	defer func() {
		dbClose, err := common.DB.DB()
		if err == nil{
			dbClose.Close()
		}
	}()
	// get an engine instance
	r := gin.Default()
	CollectRoute(r)
	r.Run()
}
