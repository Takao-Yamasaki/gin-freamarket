package infra

import (
	"log"

	"github.com/joho/godotenv"
)

func Initialize() {
	// 省略した場合は、デフォルトの.envファイルが読み込まれる
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
