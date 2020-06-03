package store

import "gitlab.srconnect.io/acuevas/graphql-server/model"

func (s *Store) CreatePatient(patient *model.Patient) error {
	_, err := s.DB.Exec(
		`INSERT INTO patients (id, name)
			VALUES ($1,$2)`,
		patient.ID,
		patient.Name)

	return err
}

func (s *Store) FetchPatient(patientID string) (*model.Patient, error) {
	row := s.DB.QueryRow(
		`SELECT id, name
		 FROM patients
		 WHERE id=$1`, patientID)

	var id, name string
	err := row.Scan(&id, &name)
	if err != nil {
		return nil, err
	}

	result := &model.Patient{
		ID:   id,
		Name: name,
	}

	return result, nil
}

func (s *Store) FetchPatients() ([]*model.Patient, error) {
	rows, err := s.DB.Query(
		`SELECT id, name
		 FROM patients`)
	if err != nil {
		return nil, err
	}

	results := make([]*model.Patient, 0)
	for rows.Next() {
		var id, name string
		if err := rows.Scan(&id, &name); err != nil {
			return nil, err
		}
		result := &model.Patient{
			ID:   id,
			Name: name,
		}

		results = append(results, result)
	}

	return results, nil
}
