package dataloader

import (
	"time"

	"gitlab.srconnect.io/acuevas/graphql-server/graph/model"
)

func (dl DataLoader) GetProviderLoader() ProviderLoader {
	return ProviderLoader{
		maxBatch: 100,
		wait:     1 * time.Millisecond,
		fetch: func(ids []string) ([]*model.Provider, []error) {
			providers, err := dl.Store.FetchProvidersWithIDs(ids)
			if err != nil {
				return nil, []error{err}
			}

			p := make(map[string]*model.Provider, len(providers))
			for _, provider := range providers {
				p[provider.ID] = provider
			}

			result := make([]*model.Provider, len(ids))
			for i, id := range ids {
				result[i] = p[id]
			}

			return result, nil
		},
	}
}
