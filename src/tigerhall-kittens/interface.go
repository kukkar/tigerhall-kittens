package tigerhall

import (
	"github.com/kukkar/common-golang/pkg/utils/queryparser"
)

type Tigerhall interface {
	CreateTiger(req ReqCreateTiger) error
	ListTigers(q queryparser.QueryParamsList,
		limit, offset int) ([]ResListTiger, int, error)
	SightATiger(id string, req ReqSightOfATiger) error
	ListSigntsOfTiger(q queryparser.QueryParamsList,
		limit, offset int) (*ResListSigntsOfTiger, error)
	CreateImage(im *Image) error
}
