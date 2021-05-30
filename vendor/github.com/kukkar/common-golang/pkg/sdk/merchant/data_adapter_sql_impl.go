package merchant

import (
	"fmt"

	sqlRepo "github.com/kukkar/common-golang/pkg/factory/sql"
)

type sqlImpl struct {
	DB *sqlRepo.MySqlConnection
}

func (this sqlImpl) GetClientSecret(clientName string) (string, error) {
	var output internalClientTable
	query := this.DB.Conn.Table(TABLE_INTERNAL_SECRET)
	query = query.Select("secret")
	query = query.Where("client_name=? ",
		clientName).Find(&output)
	if query.Error != nil {
		return "", query.Error
	}
	if output.Secret == "" {
		return "", fmt.Errorf("No secret key found")
	}
	return output.Secret, nil
}

func (this sqlImpl) GetMasterKey() (string, error) {
	var output masterKeysTable
	query := this.DB.Conn.Table(TABLE_MASTER_KEY)
	query = query.Select("secret")
	query = query.Where("active=? ",
		1).Find(&output)
	if query.Error != nil {
		return "", query.Error
	}
	if output.Secret == "" {
		return "", fmt.Errorf("No secret key found")
	}
	return output.Secret, nil
}

func (this sqlImpl) GetTokenInfo(token string) (*AppInfo, error) {
	return nil, fmt.Errorf("todo ")
}

func (this sqlImpl) GetTempTokenInfo(token string) (*AppInfo, error) {
	return nil, fmt.Errorf("todo ")
}

func (this sqlImpl) GetClientPermissions(clientName string) (*InternalCilentPermissionsTable, error) {
	var output InternalCilentPermissionsTable
	query := this.DB.Conn.Table(TableInternalClientPermissions)
	query = query.Select("*")
	query = query.Where("client_name =?",
		clientName).Find(&output)
	if query.Error != nil {
		return nil, query.Error
	}
	return &output, nil
}
