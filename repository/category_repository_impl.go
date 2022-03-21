package repository

type CategoryRepository struct {

}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category doamin.Category) domain.Category {
	sql := "insert into customer(name) values (?)"
	result, err := tx.ExecContext(ctx, SQL, Category.Name)
	helper.PanicIfError(err)
	}
	
	id, err := result.LastInsertId()
	helper.PanicIfError()
	}

	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	panic("implement me")
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	panic("implement me")
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) domain.Category {
	panic("implement me")
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, category domain.Category) []domain.Category {
	panic("implement me")
}