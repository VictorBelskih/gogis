package service

import "github.com/VictorBelskih/gogis/pkg/repository"

type Authorization interface {
}

type FarmList interface {
}

type FarmField interface {
}

type Service struct {
	Authorization
	FarmList
	FarmField
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
