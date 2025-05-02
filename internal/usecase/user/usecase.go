package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"user-service/internal/repository"

	"google.golang.org/grpc/internal/metadata"

	"github.com/golang-jwt/jwt/v5"
)

// Список ошибок
var (
	// ErrUserNotFound - ошибка, когда пользователь не найден
	ErrUserNotFound = errors.New("user not found")
	// ErrProductNotFound - ошибка, когда продукт не найден
	ErrProductNotFound = errors.New("product not found")
	// ErrPreferenceNotFound - ошибка, когда предпочтение не найдено
	ErrPreferenceNotFound = errors.New("preference not found")
	// ErrProductAlreadyExists - ошибка, когда продукт уже существует
	ErrProductAlreadyExists = errors.New("product already exists")
	// ErrMetadataNotFound - ошибка, когда метаданные не найдены
	ErrMetadataNotFound = errors.New("metadata not found")
	// ErrTokenNotFound - ошибка, когда токен не найден
	ErrTokenNotFound = errors.New("token not found")
)

var _ UserUsecase = (*user)(nil)

// UserUsecase - интерфейс для работы с пользователями
type UserUsecase interface {
	// GetUserProducts - получить список продуктов пользователя
	GetUserProducts(accessToken string) (productName map[string]string, err error)
	// GetUserPreference - получить предпочтения пользователя
	GetUserPreference(accessToken string) (preference string, err error)
	// UpdateUserPreference - обновить предпочтения пользователя
	UpdateUserPreference(accessToken string, preference string) (err error)
	// AddUserProduct - добавить продукт пользователю
	AddUserProduct(accessToken string, productName string) (err error)
	// RemoveUserProduct - удалить продукт у пользователя
	RemoveUserProduct(accessToken string, productName string) (err error)
}

type user struct {
	userRepo repository.Repository
}

// New - конструктор для создания нового экземпляра UserUsecase
func New(userRepo repository.Repository) *user {
	return &user{
		userRepo: userRepo,
	}
}

func (u *user) GetUserProducts(accessToken string) (productName map[string]string, err error) {

}

func (u *user) GetUserPreference(accessToken string) (preference string, err error) {

}

func (u *user) UpdateUserPreference(accessToken string, preference string) (err error) {

}

func (u *user) AddUserProduct(accessToken string, productName string) (err error) {

}

func (u *user) RemoveUserProduct(accessToken string, productName string) (err error) {

}

// extractUserIdFromToken - функция для извлечения id пользователя из токена
func extractUserIdFromToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", ErrMetadataNotFound
	}

	authHeader := md.Get("authorization")

	tokenString := strings.TrimPrefix(authHeader[0], "Bearer ")
	if tokenString == "" {
		return "", ErrTokenNotFound
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	userId, ok := token.Claims.(jwt.MapClaims)["user_id"].(string)

	if !ok {
		return "", ErrUserNotFound
	}

	return userId, nil
}
