package web

import (
	"context"
)
type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest)
}