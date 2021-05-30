package tigerhall

import (
	"github.com/kukkar/common-golang/pkg/utils/queryparser"
)

type Tigerhall interface {
	CreateTiger(req ReqCreateTiger) (*int, error)
	ListTigers(q queryparser.QueryParamsList,
		limit, offset int) ([]ResListTiger, error)
	SightATiger(req ReqSightOfATiger) (*int, error)
	ListSigntsOfTiger(q queryparser.QueryParamsList,
		limit, offset int) (*ResListSigntsOfTiger, error)
}
