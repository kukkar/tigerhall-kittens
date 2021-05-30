package merchant

import (
	"fmt"

	"github.com/kukkar/common-golang/pkg/logger"

	"github.com/jinzhu/copier"
	sqlRepo "github.com/kukkar/common-golang/pkg/factory/sql"
)

type oldSqlImpl struct {
	DB *sqlRepo.MySqlConnection
}

func (this oldSqlImpl) GetClientSecret(clientName string) (string, error) {
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

func (this oldSqlImpl) GetMasterKey() (string, error) {
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

func (this oldSqlImpl) GetTokenInfo(token string) (*AppInfo, error) {
	var output AppInfo
	var tableData VerifyTokenTable
	query := this.DB.Conn.Table(VerifyTable)
	query = query.Select("*")
	query = query.Where("accesstoken=? AND wipeout = ?",
		token, 0).Find(&tableData)
	if query.Error != nil {
		return nil, query.Error
	}
	logger.Logger.Info(fmt.Sprintf("sql data %v", tableData))

	copier.Copy(&output, &tableData)
	logger.Logger.Info(fmt.Sprintf("parsed data %v", output.MerchantID))
	return &output, nil
}

func (this oldSqlImpl) GetTempTokenInfo(token string) (*AppInfo, error) {
	return nil, fmt.Errorf("not contain temp table use New DB structure for that")
}

func (this oldSqlImpl) GetClientPermissions(clientName string) (*InternalCilentPermissionsTable, error) {

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
