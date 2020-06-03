package store

import (
	"fmt"

	"github.com/jmoiron/sqlx"

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
