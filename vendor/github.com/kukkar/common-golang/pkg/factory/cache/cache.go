package cache

import (
	"fmt"

	"github.com/go-redis/redis"
	concurrenthashmap "github.com/kukkar/common-golang/pkg/utils/concurrenthashmap"
)

const DEFAULT_POOL = "default"

//store cache Adapter
var cacheMap = concurrenthashmap.New()
var keyToConfigMap = make(map[string]interface{})

func InitConfigMap(configMap map[string]interface{}) {

	for key, value := range configMap {
		keyToConfigMap[key] = value
	}
}

func GetPool(key string) (*redis.Client, error) {
	if val, ok := cacheMap.Get(key); !ok {
		cache, err := getAdapter(key)
		if err != nil {
			return nil, err
		}
		cacheMap.Put(key, cache)
		return cache, nil
	} else {
		return val.(*redis.Client), nil
	}
}

func InitPool(config CacheConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr, // use default Addr
		PoolSize: config.Redis.PoolSize,
	})
	return rdb, nil
}

func getAdapter(key string) (*redis.Client, error) {
	var rdb *redis.Client

	config, err := getConfig(key)
	if err != nil {
		return nil, err

	}
	if config.Use == "redis" {
		rdb = redis.NewClient(&redis.Options{
			Addr:     config.Redis.Addr, // use default Addr
			PoolSize: config.Redis.PoolSize,
		})
	} else {
		return nil, fmt.Errorf("Selected Cache Option not implemented yet")
	}
	return rdb, nil
}

func getConfig(key string) (*CacheConfig, error) {

	var config CacheConfig
	if val, ok := keyToConfigMap[key]; ok {
		config = val.(CacheConfig)
	} else {
		return nil, fmt.Errorf("Wrong Config passed unable to assert to CacheConfig")
	}
	return &config, nil
}
