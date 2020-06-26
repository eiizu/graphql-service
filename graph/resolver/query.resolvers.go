package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"gitlab.srconnect.io/acuevas/graphql-server/graph/generated"
	"gitlab.srconnect.io/acuevas/graphql-server/model"
)

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

func (r *queryResolver) Appointment(ctx context.Context, id string) (*model.Appointment, error) {
	appointment, err := r.Store.FetchAppointment(id)
	if err != nil {
		return nil, err
	}

	return appointment, nil
}

func (r *queryResolver) Patient(ctx context.Context, id string) (*model.Patient, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Providers(ctx context.Context) ([]*model.Provider, error) {
	providers, err := r.Store.FetchProviders()
	if err != nil {
		return nil, err
	}

	return providers, nil
}

func (r *queryResolver) Provider(ctx context.Context, id string) (*model.Provider, error) {
	provider, err := r.Store.FetchProvider(id)
	if err != nil {
		return nil, err
	}

	return provider, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
