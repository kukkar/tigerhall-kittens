package merchant

import (
	"github.com/kukkar/common-golang/globalconst"
	"github.com/kukkar/common-golang/pkg/factory/cache"
	sqlRepo "github.com/kukkar/common-golang/pkg/factory/sql"
)

func GetMerchantSdk(dbToUse string) (Merchant, error) {

	conn, err := getSqlConn(dbToUse)
	if err != nil {
		return nil, err
	}
	redisCache, err := cache.GetPool(globalconst.DefaultRedisPoolKey)
	if err != nil {
		return nil, err
	}
	adapter := merchantImpl{
		dataAdapter:  conn,
		cacheAdapter: redisCache,
	}
	return adapter, nil
}

func getSqlConn(connectionKey string) (DataAdapter, error) {
	var conn DataAdapter
	switch connectionKey {
	case globalconst.OldMerchantDB:
		pool, err := sqlRepo.GetPool(globalconst.OldMerchantDB)
		if err != nil {
			return nil, err
		}
		conn = oldSqlImpl{
			DB: pool.GetConnection(),
		}

	case globalconst.DefaultDB:
		pool, err := sqlRepo.GetPool(globalconst.DefaultDB)
		if err != nil {
			return nil, err
		}
		conn = sqlImpl{
			DB: pool.GetConnection(),
		}
	}
	return conn, nil
}

func GetVerifyMerchantRedisKey(token string) string {
	return tokenRedisKey + token
}
