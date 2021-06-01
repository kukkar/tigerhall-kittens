package unittests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/kukkar/common-golang/pkg/utils/queryparser"
	appConf "github.com/kukkar/tigerhall-kittens/conf"
	"github.com/kukkar/tigerhall-kittens/src/globalconst"
	tigerhall "github.com/kukkar/tigerhall-kittens/src/tigerhall-kittens"
)

func TestCreateTiger(t *testing.T) {

	setUP()
	ctx := context.TODO()
	d, err := appConf.GetAppConfig()
	if err != nil {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", "get tiger hall instance", err.Error())
	}
	fmt.Printf("##### %v", d)
	intf, err := tigerhall.GetTigerHallKittens(ctx, tigerhall.ConfigTigerHall{
		StorageAdapter: "mongo",
	})

	if err != nil {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", "get tiger hall instance", err.Error())
	}

	err = intf.CreateTiger(mockDataCreateTiger())
	if err != nil {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", "tiger created successfully", err.Error())
	}
	t.Logf("Tiger created successfully")
}

func TestAddTigerSight(t *testing.T) {
	setUP()
	ctx := context.TODO()
	d, err := appConf.GetAppConfig()
	if err != nil {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", "get tiger hall instance", err.Error())
	}
	fmt.Printf("##### %v", d)
	intf, err := tigerhall.GetTigerHallKittens(ctx, tigerhall.ConfigTigerHall{
		StorageAdapter: "mongo",
	})

	if err != nil {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", "get tiger hall instance", err.Error())
	}

	err = intf.SightATiger("60b6027bfe8fcb3198d1f525", mockDataUpdateTiger())
	if err != nil {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", "tiger created successfully", err.Error())
		return
	}
	t.Logf("Tiger sight added successfully")
}

func TestListTigers(t *testing.T) {

	setUP()
	ctx := context.TODO()
	intf, err := tigerhall.GetTigerHallKittens(ctx, tigerhall.ConfigTigerHall{
		StorageAdapter: "mongo",
	})
	if err != nil {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", "get tiger hall instance", err.Error())
	}
	q, err := mockDatagetListQuery("sahil")
	if err != nil {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", "get tiger hall instance", err.Error())
	}
	tigerList, count, err := intf.ListTigers(q, 10, 0)
	if err != nil {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", "get tiger hall instance", err.Error())
	}
	t.Logf("tiger list %v count %v", tigerList, count)
}

func mockDatagetListQuery(name string) (queryparser.QueryParamsList, error) {

	q := "name.eq~" + name
	return parseQuery(q)

}

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

func setUP() {

	RegisterTestConfig()
	InitTestConfig()

}

func mockDataCreateTiger() tigerhall.ReqCreateTiger {
	layout := "2006-01-02T15:04:05.000Z"
	str := "2021-06-01T11:45:26.371Z"
	seenTime, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println(err)
	}
	return tigerhall.ReqCreateTiger{
		Name:   "testCreatetiger",
		DOB:    time.Now(),
		SeenAt: seenTime,
		Coordinates: tigerhall.Coordinates{
			Lat:  72.231,
			Long: 71.3131,
		},
	}
}

func mockDataUpdateTiger() tigerhall.ReqSightOfATiger {
	return tigerhall.ReqSightOfATiger{
		Coordinates: tigerhall.Coordinates{
			Lat:  72.232,
			Long: 71.3132,
		},
		TimeStamp: time.Now(),
		ImagePath: "/tigerhall/tiger/256_200",
	}
}
