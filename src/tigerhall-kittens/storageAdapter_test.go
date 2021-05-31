package tigerhall

import (
	"context"
	"testing"
	"time"

	"github.com/kukkar/tigerhall-kittens/src/tigerhall-kittens/unittests"
)

const (
	testStorageAdapter = "mongo"
)

func TestCreateTiger(t *testing.T) {

	ctx := context.TODO()
	setUP()
	st, err := getStorageAdapter("test", testStorageAdapter)
	if err != nil {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", "created tiger successfully", err.Error())
	}

	err = st.createNewTiger(ctx, TigerCollection{
		Name:       "test tiger",
		DOB:        time.Now(),
		LastSeenAt: time.Now(),
		LastSeenCoordinates: Coordinates{
			Lat:  72.132910,
			Long: 71.3820,
		},
	})
	if err != nil {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", "created tiger successfully", err.Error())
	}
}

func setUP() {
	//registering appconfig to global config
	unittests.RegisterTestConfig()
	//taking config into memory
	unittests.InitTestConfig()
}
