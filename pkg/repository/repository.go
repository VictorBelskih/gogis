package repository

import (
	"github.com/VictorBelskih/gogis"
	"github.com/jmoiron/sqlx"
)

// интерфейс авторизации объявление методов получения данных с бд
type Authorization interface {
	CreateUser(user gogis.User) (int, error)
	GetUsers() ([]gogis.User, error)
	GetUserByUsername(username string) (gogis.User, error)
}

type FarmList interface {
}

type Gis interface {
	GetField() (gogis.GeoJSON, error)
	GetFieldData() ([]gogis.Field, error)
	GetCult() ([]gogis.Cult, error)
	CreateCult(cult gogis.Cult) error
	DeleteCult(id int) error
}

type Repository struct {
	Authorization
	FarmList
	Gis
}

// создание репозиториев
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Gis:           NewGisPostgres(db),
	}

}
