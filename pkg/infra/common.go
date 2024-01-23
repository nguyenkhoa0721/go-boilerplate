package infra

import (
	"go-boilerplate/pkg/database/postgres"
	"go-boilerplate/pkg/uuid"
)

type Infra struct {
	Sql  *postgres.Sql
	Uuid *uuid.Uuid
}

func NewInfra(sql *postgres.Sql, uuid *uuid.Uuid) *Infra {
	return &Infra{
		Sql:  sql,
		Uuid: uuid,
	}
}
