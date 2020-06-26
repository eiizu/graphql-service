package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"gitlab.srconnect.io/acuevas/graphql-server/graph/generated"
	"gitlab.srconnect.io/acuevas/graphql-server/model"
)

func (r *patientResolver) Appointments(ctx context.Context, obj *model.Patient) ([]*model.Appointment, error) {
	appointments, err := r.Store.FetchAppointmentsByPatientID(obj.ID)
	if err != nil {
		return nil, err
	}

	return appointments, nil
}

// Patient returns generated.PatientResolver implementation.
func (r *Resolver) Patient() generated.PatientResolver { return &patientResolver{r} }

type patientResolver struct{ *Resolver }
