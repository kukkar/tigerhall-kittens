package merchant

type DataAdapter interface {
	GetClientSecret(clientName string) (string, error)
	GetMasterKey() (string, error)
	GetTokenInfo(token string) (*AppInfo, error)
	GetTempTokenInfo(token string) (*AppInfo, error)
	GetClientPermissions(clientName string) (*InternalCilentPermissionsTable, error)
}
