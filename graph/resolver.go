package graph

import (
	"gitlab.srconnect.io/acuevas/graphql-server/graph/model"
)

//go:generate go run github.com/99designs/gqlgen

type Store interface {
	CreatePatient(model.Patient) error
	CreateAppointment(model.Appointment) error
	CreateProvider(model.Provider) error
	FetchPatient() *model.Patient
	FetchProvider() *model.Provider
	FetchAppointment() *model.Appointment
	FetchAppointments() []*model.Appointment
}

type Resolver struct {
	Store Store
}
