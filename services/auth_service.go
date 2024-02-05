package services

import (
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"

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

	// TODO: トークンの生成処理

	return &foundUser.Email, nil
}
