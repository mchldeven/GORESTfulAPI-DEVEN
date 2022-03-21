package service

import (
	"context"
	"database/sql"
	"deven/api/helper"
	"deven/api/model/web"
	"deven/api/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
}

func (servoce *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer func() {
		err := recover()
		if err != nil {
			tx.Rollback()
			panic(err)
		} else {
			tx.Commit()
		}
	}()
}

func (servoce *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	panic("implement me")
}

func (servoce *CategoryServiceImpl) Delete(ctx context.Context, categoryId int)  {
	panic("implement me")
}

func (servoce *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	panic("implement me")
}

func (servoce *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	panic("implement me")
}