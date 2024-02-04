package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"gin-fleamarket/controllers"
	"gin-fleamarket/infra"
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"
	"gin-fleamarket/services"
)

func main() {
	infra.Initialize()
	// デバッグ用
	log.Println(os.Getenv("ENV"))
	items := []models.Item{
		{Name: "商品1", Price: 1000, Description: "説明1", SoldOut: false},
		{Name: "商品2", Price: 2000, Description: "説明2", SoldOut: true},
		{Name: "商品3", Price: 3000, Description: "説明3", SoldOut: false},
	}

	itemRepository := repositories.NewItemMemoryRepository(items)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	// ginのデフォルトルーターを指定し、rに格納
	r := gin.Default()
	// ルーターにエンドポイントを追加
	// 第二引数に実行する関数そのものを渡す
	itemRouter := r.Group("/items")
	itemRouter.GET("", itemController.FindAll)
	itemRouter.GET("/:id", itemController.FindById)
	itemRouter.POST("", itemController.Create)
	itemRouter.PUT("/:id", itemController.Update)
	itemRouter.DELETE("/:id", itemController.Delete)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//8080でサーバーを起動
	r.Run("localhost:8080")
}
