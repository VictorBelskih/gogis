package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type FarmList interface {
}

type FarmField interface {
}

type Repository struct {
	Authorization
	FarmList
	FarmField
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
