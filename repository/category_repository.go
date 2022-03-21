package repository

import "context"
	   "database/sql"
		
type CategoryRepository interface {
	Save(ctx context.Context, *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, *sql.Tx, category domain.Category) 
	FindById(ctx context.Context, *sql.Tx, category Id int) domain.Category
	FindAll(ctx context.Context, *sql.Tx) []domain.Category
}