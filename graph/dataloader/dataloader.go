package dataloader

import (
	"context"
	"net/http"

	"gitlab.srconnect.io/acuevas/graphql-server/graph/model"
)

const loadersKey = "dataloaders"

type Store interface {
	FetchProvidersWithIDs([]string) ([]*model.Provider, error)
}

type DataLoader struct {
	Store Store
}

func New(s Store) DataLoader {
	return DataLoader{
		Store: s,
	}
}

type Loaders struct {
	ProviderByID ProviderLoader
}

func (dl DataLoader) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), loadersKey, &Loaders{
			ProviderByID: dl.GetProviderLoader(),
		})

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
