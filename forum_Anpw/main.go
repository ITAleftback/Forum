package main

import (
	"forum_Anpw/common"
	"forum_Anpw/middleware"
	"forum_Anpw/routes"
	"github.com/gin-gonic/gin"
)

func main()  {
	db:=common.InitDB()

	defer db.Close()

	r:=gin.Default()
	r.Use(middleware.CrosHandler())
	r=routes.CollectRoutes(r)
	panic(r.Run())


}