package data

import (
	"go-boilerplate/pkg/database/postgres"
)

type UserRepo struct {
	*postgres.BaseSqlRepo
}

func NewUserRepo(sql *postgres.Sql) *UserRepo {
	return &UserRepo{
		BaseSqlRepo: postgres.NewBaseSqlRepo(sql),
	}
}
