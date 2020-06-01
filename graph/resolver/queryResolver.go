package resolver

import (
	"context"

	"gitlab.srconnect.io/acuevas/graphql-server/graph/generated"
	"gitlab.srconnect.io/acuevas/graphql-server/graph/model"
)

type queryResolver struct{ *Resolver }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
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
