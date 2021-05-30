package merchant

const (
	USE_DATABASE_SERVICE           = "database"
	USE_HTTP_SERVICE               = "http"
	TABLE_INTERNAL_SECRET          = "internal_client"
	TABLE_MASTER_KEY               = "master_key"
	tokenRedisKey                  = "verify:"
	VerifyTable                    = "verify"
	TokenCacheExpiryTime           = 15
	TableInternalClientPermissions = "internal_client_permissions"
)
