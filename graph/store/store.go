package store

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"gitlab.srconnect.io/acuevas/graphql-server/graph/model"

	_ "github.com/lib/pq"
)

type Store struct {
	DB *sqlx.DB
}

func New() *Store {
	dbName := "graphql-server"
	dbUser := "alfredo"
	dbHost := "localhost"
	dbPort := "5432"
	dbPassword := "test123"

	db, err := sqlx.Open("postgres",
		fmt.Sprintf(
			"dbname=%s user=%s host=%s port=%v password=%s sslmode=disable",
			dbName, dbUser, dbHost, dbPort, dbPassword))
	if err != nil {
		panic("db connection could not be opened: " + err.Error())
	}

	err = verifyStore(db)
	if err != nil {
		panic("db verification failed: " + err.Error())
	}

	return &Store{
		DB: db,
	}
}

func (s *Store) CreatePatient(patient *model.Patient) error {
	_, err := s.DB.Exec(
		`INSERT INTO patients (id, name)
			VALUES ($1,$2)`,
		patient.ID,
		patient.Name)

	return err
}

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

func (s *Store) CreateProvider(provider *model.Provider) error {
	_, err := s.DB.Exec(
		`INSERT INTO providers (id, name)
			VALUES ($1,$2)`,
		provider.ID,
		provider.Name)

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

func (s *Store) FetchProvider(providerID string) (*model.Provider, error) {
	row := s.DB.QueryRow(
		`SELECT id, name
		 FROM providers
		 WHERE id=$1`, providerID)

	var id, name string
	err := row.Scan(&id, &name)
	if err != nil {
		return nil, err
	}

	result := &model.Provider{
		ID:   id,
		Name: name,
	}

	return result, nil
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

func (s *Store) FetchProviders() ([]*model.Provider, error) {
	rows, err := s.DB.Query(
		`SELECT id, name
		 FROM providers`)
	if err != nil {
		return nil, err
	}

	results := make([]*model.Provider, 0)
	for rows.Next() {
		var id, name string
		if err := rows.Scan(&id, &name); err != nil {
			return nil, err
		}
		result := &model.Provider{
			ID:   id,
			Name: name,
		}

		results = append(results, result)
	}

	return results, nil
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

func tableExists(db *sqlx.DB, name string) (bool, error) {
	tableExists := db.QueryRow("SELECT exists ( SELECT FROM information_schema.tables WHERE table_schema = 'public' AND table_name = $1)", name)
	var exists bool
	err := tableExists.Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, err
}

func verifyStore(db *sqlx.DB) error {
	exists, err := tableExists(db, "appointments")
	if err != nil {
		return err
	}
	if !exists {
		_, err := db.Exec(
			`CREATE TABLE appointments (
				id text not null,
				date text not null,
				providerId text not null,
				patientId text not null,
			primary key (id))`)
		return err
	}
	exists, err = tableExists(db, "patients")
	if err != nil {
		return err
	}
	if !exists {
		_, err := db.Exec(
			`CREATE TABLE patients (
				id text not null,
				name text not null,
			primary key (id))`)
		return err
	}
	exists, err = tableExists(db, "providers")
	if err != nil {
		return err
	}
	if !exists {
		_, err := db.Exec(
			`CREATE TABLE providers (
				id text not null,
				name text not null,
			primary key (id))`)
		return err
	}
	return nil
}
