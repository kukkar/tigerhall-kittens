package queryparser

import (
	"fmt"
	"regexp"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

//Functions to call
// Parse
// RemoveInvalid
// GetMongoQuery
//

//
// Exposed method to be used by external clients.
//
func Parse(url string) (QueryParamsList, error) {
	querySlice, err := validateURLString(url)
	if err != nil {
		return nil, err
	}
	return QueryParamsList(querySlice), nil
}

// return ([]Query, bool)
// Notes:
// - operator and url keys should be case insensitive.
func validateURLString(url string) ([]Query, error) {
	var urlStruct []Query

	//As it will using in listing api's so there will chance where user don't want to filter anything
	//so in this case simply return empty struct
	if len(url) == 0 {
		return urlStruct, nil
	}
	paramSpilt1 := strings.Split(url, "___")

	for _, val1 := range paramSpilt1 {
		if !isValidParamFormat(val1) {
			return nil, fmt.Errorf("Invalid Param Format: %s Expected Format :%s", val1, "param.operator~value")
		}
		paramSplit2 := strings.Split(val1, "~")
		paramSplit3 := strings.Split(paramSplit2[0], ".")
		if len(paramSplit2) != 2 && len(paramSplit3) != 2 {
			return nil, fmt.Errorf("Error while separating param and it's value")
		}
		ConvrtTolowerCase := strings.ToLower(paramSplit3[1])
		if _, ok := operatorMap[ConvrtTolowerCase]; !ok {
			return nil, fmt.Errorf("Invalid operator in Param : %s", val1)
		}
		var query Query
		query.Key = paramSplit3[0]
		query.Operator = paramSplit3[1]
		query.Value = paramSplit2[1]

		urlStruct = append(urlStruct, query)
	}

	return urlStruct, nil
}

func validateData(q []Query, urldb urlDbMap) (bool, error) {
	for _, val := range q {
		ok, err := urldb[val.Key].ValidationFunc(val.Value)
		if !ok || err != nil {
			return false, nil
		}
	}
	return false, nil

}

//
//Input : string
//Output : bool
//
func isValidParamFormat(param string) bool {
	var matched bool
	if strings.Contains(param, ",") {
		matched, _ = regexp.MatchString(`^[a-zA-Z][a-zA-Z0-9_]{0,}[.]{1}[a-z]{1,}~[a-zA-Z0-9][a-zA-Z0-9/\-\s,]{0,}[^,\s]$`, param)
	} else {
		matched, _ = regexp.MatchString(`^[a-zA-Z][a-zA-Z0-9_]{0,}[.]{1}[a-zA-Z]{1,}~[a-zA-Z0-9/\-\s]{1,}$`, param)
	}

	if matched == true {
		return true
	}
	return false
}

func allowedOperatorCheck(operatorSlice []string, operator string) bool {
	for _, value := range operatorSlice {
		if value == operator {
			return true
		}
	}
	return false
}

//
// input []MongoQuery
///
//
func queryGenerator(mongoStruct []MongoQuery) map[string]interface{} {
	query := make(map[string]interface{})

	for _, value := range mongoStruct {
		switch value.Operator {
		case "eq":
			query[value.Key] = value.Value
		case "lk":
			str := value.Value.(string)
			query[value.Key] = bson.RegEx{Pattern: str, Options: "i"}
		case "in":
			query[value.Key] = value.Value
		case "lt":
			query[value.Key] = value.Value
		case "lte":
			query[value.Key] = value.Value
		case "gt":
			query[value.Key] = value.Value
		case "gte":
			query[value.Key] = value.Value
		case "ne":
			query[value.Key] = value.Value

		}
	}

	filters := make(map[string]interface{})

	filters = map[string]interface{}{
		"$and": []interface{}{
			query},
	}
	return filters
}
