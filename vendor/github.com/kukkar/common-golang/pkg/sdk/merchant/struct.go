package merchant

import "time"

type internalClientTable struct {
	Secret     string `gorm:"secret"`
	ClientName int    `gorm:"client_name"`
}
type masterKeysTable struct {
	Secret string `gorm:"secret"`
	Active int    `gorm:"active"`
}

// MerchantApptoken represents merchant_apptoken table which will carry merchant tokens
type MerchantApptoken struct {
	ID              int       `gorm:"primary_key;column:id"`
	AccessToken     string    `gorm:"column:accesstoken"`
	Status          string    `gorm:"column:status"`
	MerchantID      *int      `gorm:"column:merchant_id"`
	MerchantStoreID *int      `gorm:"column:merchant_store_id"`
	MerchantUserID  *int      `gorm:"column:merchant_user_id"`
	UserIP          string    `gorm:"column:user_ip"`
	AppVersion      string    `gorm:"column:app_version"`
	DeviceID        string    `gorm:"column:device_id"`
	UserAgent       string    `gorm:"column:user_agent"`
	SignInType      string    `gorm:"column:sign_in_type"`
	Platform        string    `gorm:"column:platform"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
}

type VerifyTokenTable struct {
	ID                 int       `gorm:"primary_key;column:id"`
	Mobile             string    `gorm:"column:mobile"`
	OTP                string    `gorm:"column:otp"`
	AccessToken        string    `gorm:"column:accesstoken"`
	Assigned           string    `gorm:"column:assigned"`
	timestamp          time.Time `gorm:"column:timestamp"`
	Wipeout            string    `gorm:"column:wipeout"`
	AccessKey          string    `gorm:"column:accesskey"`
	MerchantID         *int      `gorm:"column:merchant_id"`
	MerchantStoreID    *int      `gorm:"column:merchant_store_id"`
	MerchantUserID     *int      `gorm:"column:merchant_user_id"`
	MerchantTerminalID *int      `gorm:"column:merchant_terminal_id"`
	UserIP             string    `gorm:"column:user_ip"`
	AppVersion         string    `gorm:"column:app_version"`
	DeviceID           string    `gorm:"column:device_id"`
	UserAgent          string    `gorm:"column:user_agent"`
	SignInType         string    `gorm:"column:sign_in_type"`
	Platform           string    `gorm:"column:platform"`
}

type AppInfo struct {
	ID                 int
	Mobile             string
	AccessToken        string
	Status             string
	MerchantID         *int
	MerchantStoreID    *int
	MerchantUserID     *int
	MerchantTerminalID *int
	UserIP             string
	AppVersion         string
	DeviceID           string
	UserAgent          string
	SignInType         string
	Platform           string
}

//MerchantSignUpTempToken
type MerchantSignUpTempToken struct {
	ID           int       `gorm:"column:id"`
	MobileNumber string    `gorm:"column:mobilenumber"`
	AccessToken  string    `gorm:"column:accesstoken"`
	Status       string    `gorm:"column:status"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

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

type InternalCilentPermissionsTable struct {
	Permissions interface{} `gorm:"column:permissions"`
}
