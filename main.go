package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gin-fleamarket/controllers"
	"gin-fleamarket/infra"
	"gin-fleamarket/repositories"
	"gin-fleamarket/services"
)

// ルーターの定義
func setupRouter(db *gorm.DB) *gin.Engine {
	// itemRepository := repositories.NewItemMemoryRepository(items)
	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	authRepository := repositories.NewAuthRepository(db)
	authServive := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authServive)

	// ルーターにエンドポイントを追加
	// 第二引数に実行する関数そのものを渡す
	// ginのデフォルトルーターを指定し、rに格納
	r := gin.Default()
	itemRouter := r.Group("/items")
	authRouter := r.Group("/auth")

	itemRouter.GET("", itemController.FindAll)
	itemRouter.GET("/:id", itemController.FindById)
	itemRouter.POST("", itemController.Create)
	itemRouter.PUT("/:id", itemController.Update)
	itemRouter.DELETE("/:id", itemController.Delete)

	authRouter.POST("/signup", authController.Signup)
	authRouter.POST("/login", authController.Login)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}

func main() {
	infra.Initialize()
	db := infra.SetupDB()
	r := setupRouter(db)
	// デバッグ用
	// log.Println(os.Getenv("ENV"))
	// items := []models.Item{
	// 	{Name: "商品1", Price: 1000, Description: "説明1", SoldOut: false},
	// 	{Name: "商品2", Price: 2000, Description: "説明2", SoldOut: true},
	// 	{Name: "商品3", Price: 3000, Description: "説明3", SoldOut: false},
	// }

	//8080でサーバーを起動
	r.Run("localhost:8080")
}
