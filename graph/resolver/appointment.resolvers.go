package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"gitlab.srconnect.io/acuevas/graphql-server/graph/dataloader"
	"gitlab.srconnect.io/acuevas/graphql-server/graph/generated"
	"gitlab.srconnect.io/acuevas/graphql-server/model"
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
	return dataloader.For(ctx).ProviderByID.Load(obj.Provider.ID)
}

// Appointment returns generated.AppointmentResolver implementation.
func (r *Resolver) Appointment() generated.AppointmentResolver { return &appointmentResolver{r} }

type appointmentResolver struct{ *Resolver }
