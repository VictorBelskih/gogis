package service

import (
	"github.com/VictorBelskih/gogis"
	"github.com/VictorBelskih/gogis/pkg/repository"
)

// передача интерфейсов
type GisService struct {
	repo repository.Gis
}

// создание интерфейса интерфейсов
func NewGisService(repo repository.Gis) *GisService {
	return &GisService{repo: repo}
}
func (s *GisService) GetField() (gogis.GeoJSON, error) {
	return s.repo.GetField()
}
