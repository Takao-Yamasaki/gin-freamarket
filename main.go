package main

import "github.com/gin-gonic/gin"

func main() {
	// ginのデフォルトルーターを指定し、rに格納
	r := gin.Default()
	// ルーターにエンドポイントを追加
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//8080でサーバーを起動
	r.Run("localhost:8080")
}
