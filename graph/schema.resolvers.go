package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gitlab.srconnect.io/acuevas/graphql-server/graph/generated"
	"gitlab.srconnect.io/acuevas/graphql-server/graph/model"
)

func (r *mutationResolver) CreateAppointment(ctx context.Context, input model.NewAppointment) (*model.Appointment, error) {
	appointment := &model.Appointment{
		ID:   uuid.New().String(),
		Date: time.Now().String(),
	}

	return appointment, nil
}

func (r *mutationResolver) CreatePatient(ctx context.Context, input model.NewPatient) (*model.Patient, error) {
	patient := &model.Patient{
		ID:   uuid.New().String(),
		Name: input.Name,
	}

	return patient, nil
}

func (r *mutationResolver) CreateProvider(ctx context.Context, input model.NewProvider) (*model.Provider, error) {
	provider := &model.Provider{
		ID:   uuid.New().String(),
		Name: input.Name,
	}

	return provider, nil
}

func (r *queryResolver) Appointments(ctx context.Context) ([]*model.Appointment, error) {
	return nil, nil
}

func (r *queryResolver) Patients(ctx context.Context) ([]*model.Patient, error) {
	return nil, nil
}

func (r *queryResolver) Providers(ctx context.Context) ([]*model.Provider, error) {
	return nil, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
