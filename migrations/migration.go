package main

import (
	"fmt"
	"gin-fleamarket/infra"
	"gin-fleamarket/models"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	// ItemモデルとUserモデルのマイグレーションを実行
	if err := db.AutoMigrate(&models.Item{}, &models.User{}); err != nil {
		panic("Failed to migrate database")
	}
	fmt.Println("Success to migrate database")
}
