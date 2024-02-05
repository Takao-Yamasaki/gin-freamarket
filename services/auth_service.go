package services

import (
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Signup(email string, password string) error
	Login(email string, password string) (*string, error)
}

type AuthService struct {
	repository repositories.IAuthRepositoy
}

func NewAuthService(repository repositories.IAuthRepositoy) IAuthService {
	return &AuthService{repository: repository}
}

// ユーザーの登録処理を行う関数
func (s *AuthService) Signup(email string, password string) error {
	// パスワードをハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Email:    email,
		Password: string(hashedPassword),
	}
	return s.repository.CreateUser(user)
}

// ログイン処理をする関数
func (s *AuthService) Login(email string, password string) (*string, error) {
	// emailで対象のユーザーを検索
	foundUser, err := s.repository.FindUser(email)
	if err != nil {
		return nil, err
	}

	// ハッシュ化されたパスワードと送信されたパスワードを比較
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	// JWTトークンの生成処理を行う
	token, err := CreateToken(foundUser.ID, foundUser.Email)
	if err != nil {
		return nil, err
	}

	return token, nil
}

// JWTトークンの生成処理を行う関数
func CreateToken(userId uint, email string) (*string, error) {
	// トークンの生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   userId,                           // subject/ユーザー識別子を表す
		"email": email,                            // ユーザーのemail
		"exp":   time.Now().Add(time.Hour).Unix(), // トークンの有効期限/生成から1時間
	})

	// 秘密鍵を使って署名を行う
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}
