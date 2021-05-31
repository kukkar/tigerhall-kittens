package tigerhall

import (
	"context"
	"fmt"

	mfactory "github.com/kukkar/tigerhall-kittens/src/common/factory/mongof"
)

func GetTigerHallKittens(c context.Context,
	config ConfigTigerHall) (Tigerhall, error) {

	stAdapter, err := getStorageAdapter("", config.StorageAdapter)
	if err != nil {
		return nil, err
	}
	return &tigherhall{
		stAdapter: stAdapter,
	}, nil
}

func getStorageAdapter(key string,
	stAdapter string) (storageAdapter, error) {
	switch stAdapter {
	case StInMemory:
	case StMongo:
		md, err := getMongoAdapater(key)
		if err != nil {
			return nil, err
		}
		return &mongoAdapter{
			md,
		}, nil
	}
	return nil, fmt.Errorf("wrong choice of Adapter")
}

func getMongoAdapater(key string) (*mfactory.MDB, error) {
	if key == "" {
		key = mfactory.DEFAULT_KEY
	}
	return mfactory.GetPool(key)
}
