package repository

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrProductNotFound        = errors.New("product not found")
	ErrPreferenceNotFound     = errors.New("preference not found")
	ErrPreferenceUpdateFailed = errors.New("preference update failed")
	ErrProductAlreadyExists   = errors.New("product already exists")
	ErrQueryFailed            = errors.New("query failed")
	ErrNoRows                 = errors.New("no rows in result")
	ErrAddUserFailed          = errors.New("add user failed")
)

var _ Repository = (*repository)(nil)

type Repository interface {
	// GetProducts - получить список продуктов пользователя
	GetProducts(ctx context.Context, userId string) ([]string, error)
	// GetPreference - получить предпочтения пользователя
	GetPreference(ctx context.Context, userId string) (string, error)
	// UpdatePreference - обновить предпочтения пользователя
	UpdatePreference(ctx context.Context, userId string, preferenceName string) (error)
	// RemovePreference - удалить предпочтения пользователя
	RemovePreference(ctx context.Context, userId string) (error)
	// AddProduct - добавить продукт пользователю
	AddProduct(ctx context.Context, userId string, productName string) (error)
	// RemoveProduct - удалить продукт у пользователя
	RemoveProduct(ctx context.Context, userId string, productName string) (error)
}

type repository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *repository {
	return &repository{
		db: db,
	}
}
func (r *repository) GetProducts(ctx context.Context, userId string) ([]string, error) {
	query := `SELECT product_name FROM user_products WHERE user_id = $1`
	rows, err := r.db.Query(ctx, query, userId)
	if err != nil {
		return nil, ErrQueryFailed
	}
	defer rows.Close()
	var products []string
	for rows.Next() {
		var product string
		if err := rows.Scan(&product); err != nil {
			return nil, ErrNoRows
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *repository) GetPreference(ctx context.Context, userId string) (string, error) {
	query := `SELECT preference_name FROM user_preference WHERE user_id = $1`
	var preference string
	if err := r.db.QueryRow(ctx, query, userId).Scan(&preference); err != nil {
		return "", ErrPreferenceNotFound
	}
	return preference, nil
}

func (r *repository) UpdatePreference(ctx context.Context, userId string, preferenceName string) (error) {
	query := `SELECT user_id FROM user_preference WHERE user_id = $1`
	err := r.db.QueryRow(ctx, query, userId)
	if err != nil {
		query = `INSERT INTO user_preference (user_id, preference_name) VALUES ($1, $2)`
		if _, err := r.db.Exec(ctx, query, userId, preferenceName); err != nil {
			return ErrAddUserFailed
		}
		return nil
	}
	query = `UPDATE user_preference SET preference_name = $1 WHERE user_id = $2`
	if _, err := r.db.Exec(ctx, query, preferenceName, userId); err != nil {
		return ErrPreferenceUpdateFailed
	}
	log.Println("Preference updated successfully")
	return nil
}

func (r *repository) RemovePreference(ctx context.Context, userId string) (error) {
	query := `DELETE FROM user_preference WHERE user_id = $1`
	if _, err := r.db.Exec(ctx, query, userId); err != nil {
		return ErrPreferenceNotFound
	}
	log.Println("Preference removed successfully")
	return nil
}

func (r *repository) AddProduct(ctx context.Context, userId string, productName string) (error) {
	query := `INSERT INTO user_products (user_id, product_name) VALUES ($1, $2)`
	if _, err := r.db.Exec(ctx, query, userId, productName); err != nil {
		return ErrProductAlreadyExists
	}
	log.Println("Product added successfully")
	return nil
}

func (r *repository) RemoveProduct(ctx context.Context, userId string, productName string) (error) {
	query := `DELETE FROM user_products WHERE user_id = $1 AND product_name = $2`
	if _, err := r.db.Exec(ctx, query, userId, productName); err != nil {
		return ErrProductNotFound
	}
	log.Println("Product removed successfully")
	return nil
}
