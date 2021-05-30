package clientvalidator

type RedisMerchantIDs struct {
	Mobile             string `json:"mobile"`
	MerchantID         *int   `json:"merchant_id"`
	MerchantStoreID    *int   `json:"store_id"`
	MerchantUserID     *int   `json:"merchant_user_id"`
	MerchantTerminalID *int   `json:"merchant_terminal_id"`
	AppVersion         string `json:"app_version"`
	Token              string `json:"token"`
	Verify             int    `json:"verify"`
}
