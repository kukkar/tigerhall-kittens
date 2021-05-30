package clientvalidator

type ClientValidator interface {
	GetClientSecret(clientName string) (string, error)
	GetMasterKey() (string, error)
	VerifyToken(token string) (*RedisMerchantIDs, error)
	GetClientPermissions(clientName string) (map[string]string, error)
	GetSuperKey() (string, error)
	VerifySimBinding(instalID, simdID, deviceID string) error
}
