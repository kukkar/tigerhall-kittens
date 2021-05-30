package sql

import (
	"fmt"

	concurrenthashmap "github.com/kukkar/common-golang/pkg/utils/concurrenthashmap"
)

//const SLAVE_TYPE_POOL = "Slave"

//const DEFAULT_POOL = "default"

//Store mysql pools.
var mysqlMap = concurrenthashmap.New()
var keyToConfigMap = make(map[string]interface{})

func InitConfigMap(configMap map[string]interface{}) {

	for key, value := range configMap {
		keyToConfigMap[key] = value
	}
}

// GetPool checks whether a pool with specified type and key exists.
// if so, returns the pool
// else, initiate the pool and return it
//
//@todo: need to add logger support
func GetPool(key string) (*MySqlPool, error) {
	//finalkey := pooltype + "_" + key
	if val, ok := mysqlMap.Get(key); !ok {
		//we dont have a pool by this key, initiate new pool.
		pool, err := InitPool(key)
		if err != nil {
			return nil, fmt.Errorf("Could not initiate pool for key:%s, Error:%s",
				key, err.Error())
		}
		mysqlMap.Put(key, pool)
		return pool, nil
	} else {
		return val.(*MySqlPool), nil
	}
}

// InitPool Generally not to be called explicitely, but if you are in desperate need
// do not hesitate to call.
func InitPool(key string) (*MySqlPool, error) {
	config, err := getConfig(key)
	if err != nil {
		return nil, err
	}
	pool, err := Initiate(config)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

func getConfig(pooltype string) (MysqlConfig, error) {
	var dbconfig MysqlConfig
	if val, ok := keyToConfigMap[pooltype]; ok {
		dbconfig = val.(MysqlConfig)
	} else {
		return dbconfig, fmt.Errorf("Wrong Mysql config provided in type not match")
	}
	// @todo need to add master slave supprort along with multiple database connection support
	//@todo: below condition needs to be improved when more db are added.
	return dbconfig, nil
}
