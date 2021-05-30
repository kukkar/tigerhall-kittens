package tigerhall

import "github.com/kukkar/common-golang/pkg/utils/queryparser"

type storageAdapter interface {
	createNewTiger(req TigerCollection) (*int, error)
	getTigers(q queryparser.QueryParamsList,
		limit, offset int) ([]TigerCollection, error)
}
