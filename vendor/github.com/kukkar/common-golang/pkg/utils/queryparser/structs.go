package queryparser

import (
	"fmt"
	"strconv"
	"strings"
)

type QueryParamsList []Query

func (this *QueryParamsList) RemoveInvalid(urlDbMap urlDbMap) error {

	oldQueryParamsList := *this

	newQueryParamsList := make(QueryParamsList, 0, len(oldQueryParamsList))
	for _, value := range oldQueryParamsList {
		if _, ok := urlDbMap[value.Key]; !ok {
			continue
		}
		if !allowedOperatorCheck(urlDbMap[value.Key].AllowedOperators, value.Operator) {
			return fmt.Errorf("Error : %s operator is not allowed to use for param %s", value.Operator, value.Key)
		}
		newQueryParamsList = append(newQueryParamsList, value)
	}
	this = &newQueryParamsList
	return nil
}

func (this *QueryParamsList) GetMongoQuery(m urlDbMap) (map[string]interface{}, error) {
	err := this.RemoveInvalid(m)
	if err != nil {
		return nil, err
	}

	origDataList := *this
	mongoQueryList := make([]MongoQuery, 0, len(origDataList))
	for _, v := range *this {
		// lowerCaseKey := strings.ToLower(v.Key)
		mapval := m[v.Key]
		stringData, ok := v.Value.(string)
		if !ok {
			return nil, fmt.Errorf("Error while reading data type")
		}
		var mq MongoQuery
		var err error
		switch mapval.ValueDataType {
		case "bool":
			mq.Value, err = strconv.ParseBool(stringData)
			if err != nil {
				fmt.Errorf("Error while pasring boolean value %s", stringData)
			}

		case "int":
			mq.Value, err = strconv.Atoi(stringData)
			if err != nil {
				fmt.Errorf("Error while pasring integer value %s", stringData)
			}

		case "float64":
			mq.Value, err = strconv.ParseFloat(stringData, 64)
			if err != nil {
				fmt.Errorf("Error while parsing float64 value %s", stringData)
			}

		case "string":
			mq.Value = stringData

		case "[]int":
			sliceOfValues := strings.Split(stringData, ",")
			var intArray []int
			for _, values := range sliceOfValues {
				convertedValue, err := strconv.Atoi(values)
				if err != nil {
					fmt.Errorf("Error while parsing integer array value %s", stringData)
				}
				intArray = append(intArray, convertedValue)
			}
			mq.Value = intArray

		case "[]string":
			sliceOfValues := strings.Split(stringData, ",")
			mq.Value = sliceOfValues

		default:
			//return error
			fmt.Errorf("This data type cannot be parsed :%s", mapval.ValueDataType)
		}

		mq.Key = mapval.KeyName
		mq.Operator = v.Operator

		if mapval.ValidationFunc != nil {
			_, err = mapval.ValidationFunc(mq.Value)
			if err != nil {
				return nil, err
			}
		}

		mongoQueryList = append(mongoQueryList, mq)

	}

	qurymap := queryGenerator(mongoQueryList)
	return qurymap, nil
}

func (this *QueryParamsList) GetRawQuery(m urlDbMap) ([]Query, error) {
	err := this.RemoveInvalid(m)
	if err != nil {
		return nil, err
	}

	origDataList := *this
	rawQueryList := make([]Query, 0, len(origDataList))
	for _, v := range *this {
		// lowerCaseKey := strings.ToLower(v.Key)
		mapval := m[v.Key]
		stringData, ok := v.Value.(string)
		if !ok {
			return nil, fmt.Errorf("Error while reading data type")
		}
		var mq Query
		var err error
		switch mapval.ValueDataType {
		case "bool":
			mq.Value, err = strconv.ParseBool(stringData)
			if err != nil {
				fmt.Errorf("Error while pasring boolean value %s", stringData)
			}

		case "int":
			mq.Value, err = strconv.Atoi(stringData)
			if err != nil {
				fmt.Errorf("Error while pasring integer value %s", stringData)
			}

		case "float64":
			mq.Value, err = strconv.ParseFloat(stringData, 64)
			if err != nil {
				fmt.Errorf("Error while parsing float64 value %s", stringData)
			}

		case "string":
			mq.Value = stringData

		case "[]int":
			sliceOfValues := strings.Split(stringData, ",")
			var intArray []int
			for _, values := range sliceOfValues {
				convertedValue, err := strconv.Atoi(values)
				if err != nil {
					fmt.Errorf("Error while parsing integer array value %s", stringData)
				} else {
					intArray = append(intArray, convertedValue)
				}

			}
			mq.Value = intArray

		case "[]string":
			sliceOfValues := strings.Split(stringData, ",")
			mq.Value = sliceOfValues

		default:
			//return error
			fmt.Errorf("This data type cannot be parsed :%s", mapval.ValueDataType)
		}

		mq.Key = mapval.KeyName
		mq.Operator = v.Operator

		if mapval.ValidationFunc != nil {
			_, err = mapval.ValidationFunc(mq.Value)
			if err != nil {
				return nil, err
			}
		}

		rawQueryList = append(rawQueryList, mq)

	}

	//qurymap := queryGenerator(mongoQueryList)
	return rawQueryList, nil
}

//
// GetMysqlWhereClause
//
func (this *QueryParamsList) GetMysqlQuery(m urlDbMap) (*MysqlQuery, error) {

	err := this.RemoveInvalid(m)
	if err != nil {
		return nil, err
	}
	var mysqlQuery *MysqlQuery
	origDataList := *this
	total := len(origDataList)
	for i, v := range *this {
		// lowerCaseKey := strings.ToLower(v.Key)
		if _, ok := m[v.Key]; !ok {
			return nil, fmt.Errorf(fmt.Sprintf("Not a valid key %s", v.Key))
		}
		mapval := m[v.Key]
		stringData, ok := v.Value.(string)
		if !ok {
			return nil, fmt.Errorf("Error while reading data type")
		}
		var mq Query
		var err error
		switch mapval.ValueDataType {
		case "bool":
			mq.Value, err = strconv.ParseBool(stringData)
			if err != nil {
				fmt.Errorf("Error while pasring boolean value %s", stringData)
			}

		case "int":
			mq.Value, err = strconv.Atoi(stringData)
			if err != nil {
				fmt.Errorf("Error while pasring integer value %s", stringData)
			}

		case "float64":
			mq.Value, err = strconv.ParseFloat(stringData, 64)
			if err != nil {
				fmt.Errorf("Error while parsing float64 value %s", stringData)
			}

		case "string":
			mq.Value = stringData

		case "[]int":
			var craftedValue string
			sliceOfValues := strings.Split(stringData, ",")
			for _, eachValue := range sliceOfValues {
				craftedValue += eachValue + ","
			}
			if last := len(craftedValue) - 1; last >= 0 && craftedValue[last] == ',' {
				craftedValue = craftedValue[:last]
			}
			craftedValue = "(" + craftedValue + ")"
			mq.Value = craftedValue

		case "[]string":
			var craftedValue string
			sliceOfValues := strings.Split(stringData, ",")
			for _, eachValue := range sliceOfValues {
				craftedValue += "'" + eachValue + "',"
			}
			if last := len(craftedValue) - 1; last >= 0 && craftedValue[last] == ',' {
				craftedValue = craftedValue[:last]
			}
			craftedValue = "(" + craftedValue + ")"
			mq.Value = craftedValue

		default:
			//return error
			fmt.Errorf("This data type cannot be parsed :%s", mapval.ValueDataType)
		}

		mq.Key = mapval.KeyName
		mq.Operator = v.Operator

		if mapval.ValidationFunc != nil {
			_, err = mapval.ValidationFunc(mq.Value)
			if err != nil {
				return nil, err
			}
		}
		where, err := mysqlQueryGenerator(mq, mapval.ValueDataType)
		if err != nil {
			return nil, err
		}
		if i < total-1 {
			where += " AND "
		}
		if mysqlQuery == nil {
			temp := MysqlQuery(where)
			mysqlQuery = &temp
		} else {
			temp := MysqlQuery(mysqlQuery.ToString() + where)
			mysqlQuery = &temp
		}
	}
	return mysqlQuery, nil
}

func mysqlQueryGenerator(eachQuery Query, valueDataType interface{}) (string, error) {
	var where string
	switch eachQuery.Operator {
	case "in":
		where += eachQuery.Key + " in" + eachQuery.Value.(string)
	case "eq":
		where += eachQuery.Key + " ='" + eachQuery.Value.(string) + "'"
	case "ne":
		where += eachQuery.Key + " != '" + eachQuery.Value.(string) + "'"
	case "lk":
		where += eachQuery.Key + " like " + " '%" + eachQuery.Value.(string) + "%'"
	case "gte":
		where += eachQuery.Key + " >= " + " '" +  eachQuery.Value.(string) +"'"
	case "lte":
		where += eachQuery.Key + " <= " + " '" +  eachQuery.Value.(string) +"'"
	default:
		return "", fmt.Errorf("Invalid operator to use")
	}
	return where, nil
}

type Query struct {
	//Url Value
	Key string
	//Url operator
	Operator string

	//Url value
	Value interface{}
}

//Query -> MongoQuery
type MongoQuery struct {
	//Mongo key
	Key string
	//Mongo operator
	Operator string

	//mongo value
	Value interface{}
}

//Query -> MysqlWhere query
type MysqlQuery string

type MapValue struct {
	KeyName          string
	ValueDataType    interface{}
	ValidationFunc   func(val interface{}) (bool, error)
	AllowedOperators []string
}

//make a map for this.

var operatorMap = map[string]string{
	"in":  "in",
	"lk":  "lk",
	"eq":  "eq",
	"lt":  "lt",
	"gt":  "gt",
	"lte": "lte",
	"gte": "gte",
	"ne":  "ne",
}

type urlDbMap map[string]MapValue

//
// ToString converting mysql query to string
//
func (this *MysqlQuery) ToString() string {
	if this == nil {
		return ""
	}
	return string((*this))
}
