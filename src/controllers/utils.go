package controllers

import (
	"github.com/kukkar/common-golang/pkg/utils/queryparser"
	"github.com/kukkar/tigerhall-kittens/src/globalconst"
)

func parseQuery(q string) (queryparser.QueryParamsList, error) {

	queryParamList, queryParamErr := queryparser.Parse(q)
	if queryParamErr != nil {
		return nil, queryParamErr
	}

	validateParamErr := queryParamList.RemoveInvalid(globalconst.TigerHallQueryMap)
	if validateParamErr != nil {
		return nil, validateParamErr
	}
	return queryParamList, nil
}
