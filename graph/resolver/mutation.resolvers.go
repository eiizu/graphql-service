package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gitlab.srconnect.io/acuevas/graphql-server/graph/generated"
	"gitlab.srconnect.io/acuevas/graphql-server/model"
)

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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
