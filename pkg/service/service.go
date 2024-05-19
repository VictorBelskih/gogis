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
	GetRole() ([]gogis.Role, error)
}

type Gis interface {
	GetField() (gogis.GeoJSON, error)
	GetFieldData(id int, role int) ([]gogis.Field, error)
	GetFieldByUser(id int, role int) (gogis.GeoJSON, error)
	GetCult() ([]gogis.Cult, error)
	GetCultByID(id int) (*gogis.Cult, error)
	UpdateCult(cult gogis.Cult) error
	CalculateTotalAreaByFieldType(id int, role int) (map[string]float64, error)
	CreateCult(cult gogis.Cult) error
	DeleteCult(id int) error
	CalculateAverageHumusByClass(id int, role int) ([]HumusData, error)
	CalculateRadionuclideSummary(id int, role int) (RadionuclideSummary, error)
	AvgPhosphorByClass(id int, role int) ([]NutrientData, error)
	AvgPotassiumByClass(id int, role int) ([]NutrientData, error)
	TotalArea(id int, role int) (float64, error)
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
