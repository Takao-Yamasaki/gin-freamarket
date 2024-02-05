package main

import (
	"encoding/json"
	"gin-fleamarket/infra"
	"gin-fleamarket/models"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// 全てのテストの前に呼び出される
func TestMain(m *testing.M) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalln("Error loading .env.test file")
	}

	// テストランナーの実行
	code := m.Run()

	os.Exit(code)
}

func setupTestData(db *gorm.DB) {
	// TODO: UserIDを追加すること
	items := []models.Item{
		{Name: "テストアイテム1", Price: 1000, Description: "", SoldOut: false},
		{Name: "テストアイテム2", Price: 2000, Description: "テスト2", SoldOut: true},
		{Name: "テストアイテム3", Price: 3000, Description: "テスト2", SoldOut: false},
	}

	users := []models.User{
		{Email: "test1@example.com", Password: "test1pass"},
		{Email: "test2@example.com", Password: "test2pass"},
	}

	for _, user := range users {
		db.Create(&user)
	}

	for _, item := range items {
		db.Create(&item)
	}
}

// テストの初期化
func setup() *gin.Engine {
	db := infra.SetupDB()
	db.AutoMigrate(&models.Item{}, &models.User{})

	setupTestData(db)
	router := setupRouter(db)

	return router
}

func TestFindAll(t *testing.T) {
	// テストのセットアップ
	router := setup()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/items", nil)

	// APIリクエストの実行
	router.ServeHTTP(w, req)

	// APIの結果を取得（json形式）
	var res map[string][]models.Item
	json.Unmarshal(w.Body.Bytes(), &res)

	// アサーション
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 3, len(res["data"]))
}

// TODO: 他の関数のテストを作成すること

// TODO: 異常系のテストも追加すること
