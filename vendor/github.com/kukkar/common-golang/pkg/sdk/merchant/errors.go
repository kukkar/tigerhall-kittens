package merchant

import "errors"

var REDIS_ERROR_KEY_NOT_FOUND = errors.New("redis: nil")
var MYSQL_ERROR_RECORD_NOT_FOUND = errors.New("record not found")
var UnAuthorisedUser = errors.New("user not valid")
