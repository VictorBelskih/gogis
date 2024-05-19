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
	GetRole() ([]gogis.Role, error)
}

type FarmList interface {
}

type Gis interface {
	GetField() (gogis.GeoJSON, error)
	GetFieldData(id int, role int) ([]gogis.Field, error)
	GetCult() ([]gogis.Cult, error)
	CreateCult(cult gogis.Cult) error
	GetCultByID(id int) (*gogis.Cult, error)
	DeleteCult(id int) error
	UpdateCult(cult gogis.Cult) error
	GetFieldByUser(id int, role int) (gogis.GeoJSON, error)
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
