package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gitlab.srconnect.io/acuevas/graphql-server/graph/generated"
	"gitlab.srconnect.io/acuevas/graphql-server/graph/model"
)

func (r *appointmentResolver) Patient(ctx context.Context, obj *model.Appointment) (*model.Patient, error) {
	patient, err := r.Store.FetchPatient(obj.Patient.ID)
	if err != nil {
		return nil, err
	}

	if patient == nil {
		return nil, fmt.Errorf("patient with id: %s, not found", obj.Patient.ID)
	}

	return patient, nil
}

func (r *appointmentResolver) Provider(ctx context.Context, obj *model.Appointment) (*model.Provider, error) {
	provider, err := r.Store.FetchProvider(obj.Provider.ID)
	if err != nil {
		return nil, err
	}

	if provider == nil {
		return nil, fmt.Errorf("provider with id: %s, not found", obj.Provider.ID)
	}

	return provider, nil
}

func (r *mutationResolver) CreateAppointment(ctx context.Context, input model.NewAppointment) (*model.Appointment, error) {
	patient, err := r.Store.FetchPatient(input.PatientID)
	if err != nil {
		return nil, err
	}

	if patient == nil {
		return nil, fmt.Errorf("patient with id: %s, not found", input.PatientID)
	}

	provider, err := r.Store.FetchProvider(input.ProviderID)
	if err != nil {
		return nil, err
	}

	if provider == nil {
		return nil, fmt.Errorf("provider with id: %s, not found", input.ProviderID)
	}

	appointment := &model.Appointment{
		ID:       uuid.New().String(),
		Date:     time.Now().String(),
		Patient:  patient,
		Provider: provider,
	}

	err = r.Store.CreateAppointment(appointment)
	if err != nil {
		return nil, err
	}

	return appointment, nil
}

func (r *mutationResolver) CreatePatient(ctx context.Context, input model.NewPatient) (*model.Patient, error) {
	patient := &model.Patient{
		ID:   uuid.New().String(),
		Name: input.Name,
	}

	err := r.Store.CreatePatient(patient)
	if err != nil {
		return nil, err
	}

	return patient, nil
}

func (r *mutationResolver) CreateProvider(ctx context.Context, input model.NewProvider) (*model.Provider, error) {
	provider := &model.Provider{
		ID:   uuid.New().String(),
		Name: input.Name,
	}

	err := r.Store.CreateProvider(provider)
	if err != nil {
		return nil, err
	}

	return provider, nil
}

func (r *queryResolver) Appointments(ctx context.Context) ([]*model.Appointment, error) {
	appointments, err := r.Store.FetchAppointments()
	if err != nil {
		return nil, err
	}

	return appointments, nil
}

func (r *queryResolver) Patients(ctx context.Context) ([]*model.Patient, error) {
	patients, err := r.Store.FetchPatients()
	if err != nil {
		return nil, err
	}

	return patients, nil
}

func (r *queryResolver) Providers(ctx context.Context) ([]*model.Provider, error) {
	providers, err := r.Store.FetchProviders()
	if err != nil {
		return nil, err
	}

	return providers, nil
}

// Appointment returns generated.AppointmentResolver implementation.
func (r *Resolver) Appointment() generated.AppointmentResolver { return &appointmentResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type appointmentResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
