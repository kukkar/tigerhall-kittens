package merchant

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/kukkar/common-golang/pkg/logger"
)

type merchantImpl struct {
	dataAdapter  DataAdapter
	cacheAdapter *redis.Client
}

func (this merchantImpl) GetClientSecret(clientName string) (string, error) {
	return this.dataAdapter.GetClientSecret(clientName)
}

func (this merchantImpl) GetMasterKey() (string, error) {
	return this.dataAdapter.GetMasterKey()
}

func (this merchantImpl) GetClientPermissions(clientName string) (map[string]string, error) {
	var output map[string]string
	var outputPermissions map[string]interface{}
	data, err := this.dataAdapter.GetClientPermissions(clientName)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(fmt.Sprintf("%s", data.Permissions)), &outputPermissions)
	if err != nil {
		return nil, err
	}
	dataBytes, err := json.Marshal(outputPermissions["permissions"])
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(dataBytes, &output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (this merchantImpl) VerifyToken(token string) (*RedisMerchantIDs, error) {

	var merchantInfo *RedisMerchantIDs
	var err error
	merchantInfo, err = this.validateTokenFromCache(token)

	if err != nil {
		if err.Error() != REDIS_ERROR_KEY_NOT_FOUND.Error() {
			return nil, err
		}
	}
	if merchantInfo != nil {
		return merchantInfo, nil
	}
	//need to validate in mysql db as well to sure user token is not valid
	tokenInfo, err := this.dataAdapter.GetTokenInfo(token)
	if err != nil {
		if err.Error() == MYSQL_ERROR_RECORD_NOT_FOUND.Error() {
			return nil, UnAuthorisedUser
		}
		return nil, err
	}
	if tokenInfo == nil {
		return nil, UnAuthorisedUser
	}

	merchantInfo = &RedisMerchantIDs{
		Mobile:             tokenInfo.Mobile,
		MerchantID:         tokenInfo.MerchantID,
		MerchantStoreID:    tokenInfo.MerchantStoreID,
		MerchantUserID:     tokenInfo.MerchantUserID,
		MerchantTerminalID: tokenInfo.MerchantTerminalID,
		Token:              token,
		AppVersion:         tokenInfo.AppVersion,
		Verify:             1,
	}
	logger.Logger.Info(fmt.Sprintf("info %v", tokenInfo))
	logger.Logger.Info(fmt.Sprintf("info %v", merchantInfo))
	err = this.setTokenInfo(token, (*merchantInfo))
	if err != nil {
		return nil, err
	}
	return merchantInfo, nil
}

func (this merchantImpl) validateTokenFromCache(token string) (*RedisMerchantIDs, error) {
	var output RedisMerchantIDs
	redisKey := GetVerifyMerchantRedisKey(token)
	cmd := this.cacheAdapter.Get(redisKey)
	result, err := cmd.Result()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(result), &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

func (this merchantImpl) setTokenInfo(token string, data RedisMerchantIDs) error {

	cacheDataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	logger.Logger.Info(fmt.Sprintf("data %v", data))
	cmd := this.cacheAdapter.Set(GetVerifyMerchantRedisKey(token), cacheDataBytes, time.Duration(TokenCacheExpiryTime)*time.Minute)
	_, err = cmd.Result()
	if err != nil {
		return err
	}
	return nil
}
