package merchant

import "fmt"

type merchantServiceImpl struct {
}

// @todo
func (this merchantServiceImpl) GetClientSecret(clientName string) (string, error) {
	return "DzGxcbyadBQ0ecUqYUZqqFxhCqEfCXGJrXxaNjWNUupRxNBsFuk2xdL7/r6D0z6B", nil

}

func (this merchantServiceImpl) GetMasterKey() (string, error) {
	return "ZARR31nBtrzxrdkUojAwt0UPStax8MhcpzciFAjGg4KXBL17SXw+K3slip5+qVpn", nil
}

func (this merchantServiceImpl) GetTokenInfo(token string) (*MerchantApptoken, error) {
	return nil, nil
}

func (this merchantServiceImpl) GetTempTokenInfo(token string) (*MerchantSignUpTempToken, error) {
	return nil, nil
}
func (this merchantServiceImpl) GetClientPermissions(clientName string) (*InternalCilentPermissionsTable, error) {
	return nil, fmt.Errorf("todo ")
}
