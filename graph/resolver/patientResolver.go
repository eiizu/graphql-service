package resolver

import (
	"context"

	"gitlab.srconnect.io/acuevas/graphql-server/graph/generated"
	"gitlab.srconnect.io/acuevas/graphql-server/graph/model"
)

type patientResolver struct{ *Resolver }

// Appointment returns generated.PatientResolver implementation.
func (r *Resolver) Patient() generated.PatientResolver {
	return &patientResolver{r}
}

func (r *patientResolver) Appointments(ctx context.Context, obj *model.Patient) ([]*model.Appointment, error) {
	appointments, err := r.Store.FetchAppointmentsByPatientID(obj.ID)
	if err != nil {
		return nil, err
	}

	return appointments, nil
}
