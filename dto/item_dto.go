package dto

// DTO用の構造体
// requiredでバリデーションを追加
type CreateItemInput struct {
	Name        string `json:"name" binding:"required,min=2"`
	Price       uint   `json:"price" binding:"required,min=1,max=999999"`
	Description string `json:"description"`
}
