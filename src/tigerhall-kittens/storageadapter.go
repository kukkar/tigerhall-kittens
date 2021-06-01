package tigerhall

import (
	"context"

	"github.com/kukkar/common-golang/pkg/utils/queryparser"
)

type storageAdapter interface {
	createNewTiger(ctx context.Context,
		req TigerCollection) error
	getTigers(ctx context.Context, q queryparser.QueryParamsList,
		limit, page int, sortBy string, sortOrder string) ([]TigerCollection, int, error)
	addTigerSight(ctx context.Context, id string, req SightData) error

	getTigerData(ctx context.Context,
		id string) (*TigerCollection, error)
}
