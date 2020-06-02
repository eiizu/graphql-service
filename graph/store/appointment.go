package store

import "gitlab.srconnect.io/acuevas/graphql-server/graph/model"

func (s *Store) CreateAppointment(appointment *model.Appointment) error {
	_, err := s.DB.Exec(
		`INSERT INTO appointments (id, date, patientId, providerId)
			VALUES ($1,$2,$3,$4)`,
		appointment.ID,
		appointment.Date,
		appointment.Patient.ID,
		appointment.Provider.ID)

	return err
}

func (s *Store) FetchAppointment(appointmentID string) (*model.Appointment, error) {
	row := s.DB.QueryRow(
		`SELECT id, date, patientId, providerId
		 FROM appointments
		 WHERE id=$1`, appointmentID)

	var id, date, patientID, providerID string
	err := row.Scan(&id, &date, &patientID, &providerID)
	if err != nil {
		return nil, err
	}
	result := &model.Appointment{
		ID:   id,
		Date: date,
		Patient: &model.Patient{
			ID: patientID,
		},
		Provider: &model.Provider{
			ID: providerID,
		},
	}

	return result, nil
}

func (s *Store) FetchAppointments() ([]*model.Appointment, error) {
	rows, err := s.DB.Query(
		`SELECT id, date, patientId, providerId
		 FROM appointments`)
	if err != nil {
		return nil, err
	}

	results := make([]*model.Appointment, 0)
	for rows.Next() {
		var id, date, patientID, providerID string
		if err := rows.Scan(&id, &date, &patientID, &providerID); err != nil {
			return nil, err
		}
		result := &model.Appointment{
			ID:   id,
			Date: date,
			Patient: &model.Patient{
				ID: patientID,
			},
			Provider: &model.Provider{
				ID: providerID,
			},
		}

		results = append(results, result)
	}

	return results, nil
}

func (s *Store) FetchAppointmentsByPatientID(id string) ([]*model.Appointment, error) {
	rows, err := s.DB.Query(
		`SELECT id, date, patientId, providerId
		 FROM appointments
		 WHERE patientId = $1`, id)
	if err != nil {
		return nil, err
	}

	results := make([]*model.Appointment, 0)
	for rows.Next() {
		var id, date, patientID, providerID string
		if err := rows.Scan(&id, &date, &patientID, &providerID); err != nil {
			return nil, err
		}
		result := &model.Appointment{
			ID:   id,
			Date: date,
			Patient: &model.Patient{
				ID: patientID,
			},
			Provider: &model.Provider{
				ID: providerID,
			},
		}

		results = append(results, result)
	}

	return results, nil
}
