package tigerhall

import (
	"fmt"

	"github.com/kukkar/common-golang/pkg/utils/queryparser"
)

type mongoAdapter struct {
}

func (this *mongoAdapter) createNewTiger(req TigerCollection) (*int, error) {
	return nil, fmt.Errorf("todo")
}

func (this *mongoAdapter) getTigers(q queryparser.QueryParamsList,
	limit, offset int) ([]TigerCollection, error) {
	return nil, fmt.Errorf("todo")
}
