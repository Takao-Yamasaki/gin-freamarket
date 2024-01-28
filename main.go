package main

import (
	"github.com/gin-gonic/gin"

	"gin-fleamarket/controllers"
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"
	"gin-fleamarket/services"
)

func main() {
	items := []models.Item{
		{ID: 1, Name: "商品1", Price: 1000, Description: "説明1", SoldOut: false},
		{ID: 2, Name: "商品2", Price: 2000, Description: "説明2", SoldOut: true},
		{ID: 3, Name: "商品3", Price: 3000, Description: "説明3", SoldOut: false},
	}

	itemRepository := repositories.NewItemMemoryRepository(items)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	// ginのデフォルトルーターを指定し、rに格納
	r := gin.Default()
	// ルーターにエンドポイントを追加
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 第二引数に実行する関数そのものを渡す
	r.GET("/items", itemController.FindAll)
	r.GET("/items/:id", itemController.FindById)
	r.POST("/items", itemController.Create)

	//8080でサーバーを起動
	r.Run("localhost:8080")
}
