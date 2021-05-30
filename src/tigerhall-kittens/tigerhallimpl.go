package tigerhall

import (
	"fmt"

	"github.com/kukkar/common-golang/pkg/utils/queryparser"
)

type tigherhall struct {
	stAdapter storageAdapter
}

func (this *tigherhall) CreateTiger(req ReqCreateTiger) (*int, error) {

	return nil, fmt.Errorf("todo")
}

func (this *tigherhall) ListTigers(q queryparser.QueryParamsList,
	limit, offset int) ([]ResListTiger, error) {

	return nil, fmt.Errorf("todo")
}

func (this *tigherhall) SightATiger(req ReqSightOfATiger) (*int, error) {
	return nil, fmt.Errorf("todo")
}

func (this *tigherhall) ListSigntsOfTiger(q queryparser.QueryParamsList,
	limit, offset int) (*ResListSigntsOfTiger, error) {

	return nil, fmt.Errorf("todo")
}
