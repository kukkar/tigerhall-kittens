package merchant

type Merchant interface {
	GetClientSecret(clientName string) (string, error)
	GetMasterKey() (string, error)
	VerifyToken(token string) (*RedisMerchantIDs, error)
	GetClientPermissions(clientName string) (map[string]string, error)
}
