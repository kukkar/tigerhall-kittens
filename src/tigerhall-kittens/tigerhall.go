package tigerhall

import (
	"context"
	"fmt"

	appConf "github.com/kukkar/tigerhall-kittens/conf"
	mfactory "github.com/kukkar/tigerhall-kittens/src/common/factory/mongof"
)

func GetTigerHallKittens(c context.Context,
	config ConfigTigerHall) (Tigerhall, error) {

	stAdapter, err := getStorageAdapter(mfactory.DEFAULT_KEY, config.StorageAdapter)
	if err != nil {
		return nil, err
	}
	imStAdapter, err := GetImageAdapter()
	if err != nil {
		return nil, err
	}
	return &tigherhall{
		stAdapter:      stAdapter,
		imageStAdapter: imStAdapter,
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

func GetImageAdapter() (ImageStorageAdapter, error) {

	config, _ := appConf.GetAppConfig()

	adapter2use := config.ImageStorage.Use
	switch adapter2use {
	case ADAPTER_TYPE_LOCAL:
		return InitializeLocalStorage(
			LocalStorageConf{
				Path: config.ImageStorage.Local.Path,
			})
	default:
		return nil, fmt.Errorf("Wrong Adapter supplied for Image storage")
	}
}
