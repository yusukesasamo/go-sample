package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yusukesasamo/go-sample/src/controller"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/go-sample/api/v1")
	{
		v1.GET("/user", controller.UsersGET)
		v1.POST("/user", controller.UserPOST)
		v1.PATCH("/user/:id", controller.UserPATCH)
		v1.DELETE("/user/:id", controller.UserDELETE)
		v1.GET("/userAuth", controller.userAuth)
		v1.GET("/userPurchaseHistory", controller.UserPurchaseHistoriesGET)
		v1.POST("/userPurchaseHistory", controller.UserPurchaseHistoryPOST)
		v1.GET("/item", controller.ItemsGET)
		v1.POST("/item", controller.ItemPOST)
		v1.PATCH("/item/:id", controller.ItemPATCH)
		v1.DELETE("/item/:id", controller.ItemDELETE)
	}
	router.Run(":9000")
}
