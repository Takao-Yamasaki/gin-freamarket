package dto

// DTO用の構造体
// requiredでバリデーションを追加
type CreateItemInput struct {
	Name        string `json:"name" binding:"required,min=2"`
	Price       uint   `json:"price" binding:"required,min=1,max=999999"`
	Description string `json:"description"`
}

// フィールドがない場合は更新したくないので、ポインタ型
type UpdateItemInput struct {
	Name        *string `json:"name" binding:"omitnil,min=2"`
	Price       *uint   `json:"price" binding:"omitnil,min=1,max=999999"`
	Description *string `json:"description"`
	SoldOut     *bool   `json:"soldOut"`
}
