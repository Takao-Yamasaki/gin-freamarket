package repositories

import "gin-fleamarket/models"

// インターフェースの定義
// リポジトリが満たすべきメソッドの定義を記述
type IItemRepository interface {
	FindAll() (*[]models.Item, error)
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

func (r *ItemMemoryRepository) FindAll() (*[]models.Item, error) {
	return &r.items, nil
}
