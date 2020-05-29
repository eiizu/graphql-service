package graph

import (
	"gitlab.srconnect.io/acuevas/graphql-server/graph/model"
)

//go:generate go run github.com/99designs/gqlgen

type Store interface {
	CreatePatient(*model.Patient) error
	CreateAppointment(*model.Appointment) error
	CreateProvider(*model.Provider) error
	FetchPatient(string) (*model.Patient, error)
	FetchProvider(string) (*model.Provider, error)
	FetchPatients() ([]*model.Patient, error)
	FetchProviders() ([]*model.Provider, error)
	FetchAppointments() ([]*model.Appointment, error)
}

type Resolver struct {
	Store Store
}
