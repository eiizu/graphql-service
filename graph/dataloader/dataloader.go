package dataloader

import (
	"context"
	"net/http"
	"time"

	"gitlab.srconnect.io/acuevas/graphql-server/graph/model"
)

const providerLoaderKey = "providerLoaderKey"

type Store interface {
	FetchProvidersWithIDs([]string) ([]*model.Provider, error)
}

func Middleware(s Store, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		providerLoader := ProviderLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(ids []string) ([]*model.Provider, []error) {
				providers, err := s.FetchProvidersWithIDs(ids)

				return providers, []error{err}
			},
		}

		ctx := context.WithValue(r.Context(), providerLoaderKey, &providerLoader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetProviderLoader(ctx context.Context) *ProviderLoader {
	return ctx.Value(providerLoaderKey).(*ProviderLoader)
}
