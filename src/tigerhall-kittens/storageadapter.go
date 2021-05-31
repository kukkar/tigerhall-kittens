package tigerhall

import (
	"context"

	"github.com/kukkar/common-golang/pkg/utils/queryparser"
)

type storageAdapter interface {
	createNewTiger(ctx context.Context,
		req TigerCollection) error
	getTigers(ctx context.Context, q queryparser.QueryParamsList,
		limit, offset int) ([]TigerCollection, error)
}
