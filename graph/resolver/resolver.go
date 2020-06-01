package resolver

import (
	"gitlab.srconnect.io/acuevas/graphql-server/graph/model"
)

//go:generate go run github.com/99designs/gqlgen

type Store interface {
	CreatePatient(*model.Patient) error
	CreateAppointment(*model.Appointment) error
	CreateProvider(*model.Provider) error
	FetchPatient(string) (*model.Patient, error)
	FetchPatients() ([]*model.Patient, error)
	FetchProviders() ([]*model.Provider, error)
	FetchProvider(string) (*model.Provider, error)
	FetchAppointment(string) (*model.Appointment, error)
	FetchAppointments() ([]*model.Appointment, error)
}

type Resolver struct {
	Store Store
}

func New(store Store) *Resolver {
	return &Resolver{
		Store: store,
	}
}
