package dto

// DTO用の構造体
type SignupInput struct {
	// 必須項目かつemail形式
	Email    string `json:"email" binging:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}
