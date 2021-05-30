package otp

import (
	"fmt"

	"github.com/kukkar/common-golang/pkg/utils/clientvalidator"
)

//OTPConfig
type OTPConfig struct {
	IPPort          string
	Version         string
	SuperKey        string
	MerchantDBToUse string
	ClientValidator clientvalidator.ClientValidator
}

func (this *OTPConfig) validateConfig() error {
	if this.IPPort == "" {
		return fmt.Errorf("IPPort can not be empty")
	}
	if this.SuperKey == "" {
		return fmt.Errorf("Super Key can not be empty")
	}
	return nil
}

type verifyOTPReq struct {
	UUID string `json:"uuid"`
	OTP  string `json:"otp"`
}

type verifyOTPRes struct {
	Status string `json:"status"`
}

type GenerateOTPReq struct {
	Mobile      string `json:"phoneNumber"`
	Email       string `json:"email"`
	Source      string `json:"source"`
	Purpose     string `json:"purpose"`
	ClientName  string `json:"clientName"`
	SMSTemplate string `json:"smsTemplate"`
}

type ResendOTPReq struct {
	UUID        string `json:"uuid"`
	ClientName  string `json:"clientName"`
	SMSTemplate string `json:"smsTemplate"`
}

type generateOTPServiceResponse struct {
	Data   generateOTPData `json:"data"`
	Status string          `json:"status"`
}
type generateOTPData struct {
	MobileUUID string `json:"mobileuuid"`
	EmailUUID  string `json:"emailuuid"`
}

type GenerateOTPResponse struct {
	MobileUUID string
	EmailUUID  string
}

type ResendOTPServiceRes struct {
	Data   ResendOTPServiceData `json:"data"`
	Status string               `json:"status"`
}

type ResendOTPServiceData struct {
	Resend bool `json:"resend"`
}

type VerifyOtpResV2 struct {
	Data   VerifyOtpData `json:"data"`
	Status string        `json:"status"`
}

type VerifyOtpData struct {
	UserMobile string `json:"userMobile"`
	UserEmail  string `json:"userEmail"`
	Verified   bool   `json:"verified"`
}
