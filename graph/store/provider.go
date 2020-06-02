package store

import (
	"github.com/jmoiron/sqlx"
	"gitlab.srconnect.io/acuevas/graphql-server/graph/model"
)

func (s *Store) CreateProvider(provider *model.Provider) error {
	_, err := s.DB.Exec(
		`INSERT INTO providers (id, name)
			VALUES ($1,$2)`,
		provider.ID,
		provider.Name)

	return err
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

func (s *Store) FetchProvidersWithIDs(ids []string) ([]*model.Provider, error) {
	q, args, err := sqlx.In("SELECT id, name FROM providers WHERE id IN (?)", ids)
	if err != nil {
		return nil, err
	}

	q = sqlx.Rebind(sqlx.DOLLAR, q)
	rows, err := s.DB.Query(q, args...)
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
