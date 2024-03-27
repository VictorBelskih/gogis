package service

import (
	"github.com/VictorBelskih/gogis"
	"github.com/VictorBelskih/gogis/pkg/repository"
)

// интерфейсы для обработки данных
type Authorization interface {
	CreateUser(user gogis.User) (int, error)
	GetUsers() ([]gogis.User, error)
	AuthenticateUser(username, password string) (string, error)
	ParseJWTToken(tokenString string) (gogis.User, error)
}

type Gis interface {
	GetField() (gogis.GeoJSON, error)
}

type FarmField interface {
}

type Service struct {
	Authorization
	Gis
	FarmField
}

// создания интерфейсов для обработки данных
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Gis:           NewGisService(repos.Gis),
	}
}
