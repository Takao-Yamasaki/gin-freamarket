package repositories

import (
	"errors"
	"gin-fleamarket/models"

	"gorm.io/gorm"
)

// インターフェースの定義
// リポジトリが満たすべきメソッドの定義を記述
type IItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint, userId uint) (*models.Item, error)
	Create(newItem models.Item) (*models.Item, error)
	Update(updateItem models.Item) (*models.Item, error)
	Delete(itemId uint, useId uint) error
}

// 上記インターフェースを満たす具体的な実装
type ItemMemoryRepository struct {
	items []models.Item
}

// ItemMemoryRepositoryを生成するためのFactory関数
// 具体的な実装がインターフェースの定義を満たしていない場合にエラーとなる
func NewItemMemoryRepository(items []models.Item) IItemRepository {
	return &ItemMemoryRepository{items: items}
}

// 商品の一覧を取得するメソッド
func (r *ItemMemoryRepository) FindAll() (*[]models.Item, error) {
	return &r.items, nil
}

// 商品IDによる検索を行うメソッド
func (r *ItemMemoryRepository) FindById(itemId uint, userId uint) (*models.Item, error) {
	for _, item := range r.items {
		if item.ID == itemId {
			return &item, nil
		}
	}
	return nil, errors.New("item not found")
}

// 商品を新規登録するメソッド
func (r *ItemMemoryRepository) Create(newItem models.Item) (*models.Item, error) {
	newItem.ID = uint(len(r.items) + 1)
	r.items = append(r.items, newItem)
	return &newItem, nil
}

// 商品情報を更新するメソッド
func (r *ItemMemoryRepository) Update(updateItem models.Item) (*models.Item, error) {
	for i, v := range r.items {
		if v.ID == updateItem.ID {
			r.items[i] = updateItem
			return &r.items[i], nil
		}
	}
	return nil, errors.New("unexpected error")
}

// 商品情報を削除するメソッド
func (r *ItemMemoryRepository) Delete(itemId uint, userId uint) error {
	for i, v := range r.items {
		if v.ID == itemId && v.UserID == userId {
			r.items = append(r.items[:i], r.items[i+1:]...)
			return nil
		}
	}
	return errors.New("item not found")
}

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) IItemRepository {
	return &ItemRepository{db: db}
}

func (r *ItemRepository) Create(newItem models.Item) (*models.Item, error) {
	result := r.db.Create(&newItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newItem, nil
}

func (r *ItemRepository) Delete(itemId uint, userId uint) error {
	// 削除対象の商品を検索
	deletelItem, err := r.FindById(itemId, userId)
	if err != nil {
		return err
	}

	// 削除対象の商品を論理削除
	result := r.db.Delete(deletelItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ItemRepository) FindAll() (*[]models.Item, error) {
	var items []models.Item
	result := r.db.Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return &items, nil
}

func (r *ItemRepository) FindById(itemId uint, userId uint) (*models.Item, error) {
	var item models.Item
	result := r.db.First(&item, "id = ? AND user_id = ?", itemId, userId)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("item not found")
		}
		return nil, result.Error
	}
	return &item, nil
}

func (r *ItemRepository) Update(updateItem models.Item) (*models.Item, error) {
	result := r.db.Save(&updateItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updateItem, nil
}
