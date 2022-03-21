package repository

import (
	"context"
	"database/sql"
	"deven/api/model/domain"
)
	   
		
type CategoryRepository interface {
	Save(ctx context.Context, *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, *sql.Tx, category domain.Category) domain.Category
	FindById(ctx context.Context, *sql.Tx, category Id int) domain.Category
	FindAll(ctx context.Context, *sql.Tx) []domain.Category
}