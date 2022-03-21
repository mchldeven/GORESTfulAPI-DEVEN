package service

import "context"

type CategoryService interface {
	Create(ctx context.Context)
	Update(ctx context.Context)
	Delete(ctx context.Context)
	FindById(ctx context.Context)
	FindAll(ctx context.Context)
}
