package service

import "database/sql"
	   "deven/api/repository"

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
}

