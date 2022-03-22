package helper

import (
	"deven/api/model/domain"
	"deven/api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse {
		Id: category.Id,
		Name: category.Name, 
	}
}