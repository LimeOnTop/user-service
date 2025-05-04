package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"user-service/internal/adapter/token"
	"user-service/internal/repository"
)

// Список ошибок
var (
	// ErrInvalidToken - ошибка, когда токен недействителен
	ErrInvalidToken = errors.New("invalid token")
)

var _ UserUseCase = (*user)(nil)

// UserUsecase - интерфейс для работы с пользователями
type UserUseCase interface {
	// GetUserProducts - получить список продуктов пользователя
	GetUserProducts(accessToken string) (products []string, err error)
	// GetUserPreference - получить предпочтения пользователя
	GetUserPreference(accessToken string) (preferenceName string, err error)
	// UpdateUserPreference - обновить предпочтения пользователя
	UpdateUserPreference(accessToken string, preferenceName string) (err error)
	// RemoveUserPreference - удалить предпочтения пользователя
	RemoveUserPreference(accessToken string) (err error)
	// AddUserProduct - добавить продукт пользователю
	AddUserProduct(accessToken string, productName string) (err error)
	// RemoveUserProduct - удалить продукт у пользователя
	RemoveUserProduct(accessToken string, productName string) (err error)
}

type user struct {
	userRepo     repository.Repository
	tokenService token.Token
}

// New - конструктор для создания нового экземпляра UserUsecase
func New(userRepo repository.Repository, tokenService token.Token) *user {
	return &user{
		userRepo:     userRepo,
		tokenService: tokenService,
	}
}

func (u *user) GetUserProducts(accessToken string) (products []string, err error) {
	userId, err := u.extractUserIdFromToken(accessToken)
	if err != nil {
		return nil, err
	}

	products, err = u.userRepo.GetProducts(context.Background(), userId.String())
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (u *user) GetUserPreference(accessToken string) (preferenceName string, err error) {
	userId, err := u.extractUserIdFromToken(accessToken)
	if err != nil {
		return "", err
	}

	preferenceName, err = u.userRepo.GetPreference(context.Background(), userId.String())
	if err != nil {
		return "", err
	}

	return preferenceName, nil
}

func (u *user) UpdateUserPreference(accessToken string, preferenceName string) (err error) {
	userId, err := u.extractUserIdFromToken(accessToken)
	if err != nil {
		return err
	}
	err = u.userRepo.UpdatePreference(context.Background(), userId.String(), preferenceName)

	return err
}

func (u *user) RemoveUserPreference(accessToken string) (err error) {
	userId, err := u.extractUserIdFromToken(accessToken)
	if err != nil {
		return err
	}
	err = u.userRepo.RemovePreference(context.Background(), userId.String())
	if err != nil {
		return err
	}

	return nil
}

func (u *user) AddUserProduct(accessToken string, productName string) (err error) {
	userId, err := u.extractUserIdFromToken(accessToken)
	if err != nil {
		return err
	}
	err = u.userRepo.AddProduct(context.Background(), userId.String(), productName)
	if err != nil {
		return err
	}

	return nil
}

func (u *user) RemoveUserProduct(accessToken string, productName string) (err error) {
	userId, err := u.extractUserIdFromToken(accessToken)
	if err != nil {
		return err
	}
	err = u.userRepo.RemoveProduct(context.Background(), userId.String(), productName)
	if err != nil {
		return err
	}

	return nil
}

// extractUserIdFromToken - функция для извлечения id пользователя из токена
func (u *user) extractUserIdFromToken(accessToken string) (uuid.UUID, error) {
	valid, userId, err := u.tokenService.ValidateToken(accessToken)

	if !valid {
		return uuid.Nil, err
	}

	if err != nil {
		return uuid.Nil, err
	}
	return userId, nil
}
