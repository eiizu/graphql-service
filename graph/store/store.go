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

	return &Store{
		DB: db,
	}
}

func (s *Store) CreatePatient(model.Patient) error {
	return nil
}

func (s *Store) CreateAppointment(model.Appointment) error {
	return nil
}

func (s *Store) CreateProvider(model.Provider) error {
	return nil
}

func (s *Store) FetchPatient() *model.Patient {
	return nil
}

func (s *Store) FetchProvider() *model.Provider {
	return nil
}

func (s *Store) FetchAppointment() *model.Appointment {
	return nil
}

func (s *Store) FetchAppointments() []*model.Appointment {
	return nil
}
