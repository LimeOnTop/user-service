package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ Repository = (*repository)(nil)
type Repository interface {
	// GetProducts - получить список продуктов пользователя
	GetProducts(ctx context.Context, userId string, ) ([]string, error)
	// GetPreference - получить предпочтения пользователя
	GetPreference(ctx context.Context, userId string) (string, err error)
	// UpdatePreference - обновить предпочтения пользователя
	UpdatePreference(ctx context.Context, userId string) (bool, error)
	// AddProduct - добавить продукт пользователю
	AddProduct(ctx context.Context, userId string) (bool, error)
	// RemoveProduct - удалить продукт у пользователя
	RemoveProduct(ctx context.Context, userId string) (bool, error)
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *repository {
	return &repository{
		db: db,
	}
}
func (r *repository) GetProducts(ctx context.Context) (map[string]string, error) {
	query := `SELECT product_name FROM user_products WHERE user_id = $1`
	_, err := r.db.Exec(ctx, query, )
}

func (r *repository) GetPreference(ctx context.Context) (preference string, err error) {

}

func (r *repository) UpdatePreference(ctx context.Context) (string, error) {

}

func (r *repository) AddProduct(ctx context.Context) (string, error) {

}

func (r *repository) RemoveProduct(ctx context.Context) (string, error) {

}
