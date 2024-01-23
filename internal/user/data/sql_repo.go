package data

import (
	"go-boilerplate/pkg/database/postgres"
	"go-boilerplate/pkg/infra"
)

type UserRepo struct {
	*infra.BaseSqlRepo
}

func NewUserRepo(sql *postgres.Sql) *UserRepo {
	return &UserRepo{
		BaseSqlRepo: infra.NewBaseSqlRepo(sql),
	}
}
