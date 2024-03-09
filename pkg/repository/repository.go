package repository

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

func NewRepository() *Repository {
	return &Repository{}
}
