package imagestorage

import (
	"fmt"

	tigerhall "github.com/kukkar/tigerhall-kittens/src/tigerhall-kittens"

	concurrenthashmap "github.com/kukkar/common-golang/pkg/utils/concurrenthashmap"
	appConf "github.com/kukkar/tigerhall-kittens/conf"
)

const DEFAULT_KEY = "default"

//Store mysql pools.
var imageStorageMap = concurrenthashmap.New()

func GetAdapter(key string) (tigerhall.ImageStorageAdapter, error) {
	finalkey := key
	if val, ok := imageStorageMap.Get(finalkey); !ok {
		adapter, err := getAdapter()
		if err != nil {
			return nil, fmt.Errorf(
				"Could not getAdapter for key:%s,Error:%s",
				key, err.Error())
		}
		imageStorageMap.Put(finalkey, adapter)
		return adapter, nil
	} else {
		return val.(tigerhall.ImageStorageAdapter), nil
	}
}

func getAdapter() (tigerhall.ImageStorageAdapter, error) {

	config, _ := appConf.GetAppConfig()

	adapter2use := config.ImageStorage.Use
	switch adapter2use {
	case tigerhall.ADAPTER_TYPE_LOCAL:
		return tigerhall.InitializeLocalStorage(
			tigerhall.LocalStorageConf{
				Path: config.ImageStorage.Local.Path,
			})
	default:
		return nil, fmt.Errorf("Wrong Adapter supplied for Image storage")
	}
}
