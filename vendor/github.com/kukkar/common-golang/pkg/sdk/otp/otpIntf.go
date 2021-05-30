package otp

type OTPSdk interface {
	VerfiyOTP(otp string, uuid string, clientName string) error
	VerfiyOTPV2(otp string, uuid string, clientName string) (*VerifyOtpData, error)
	GenerateOTP(req GenerateOTPReq) (*GenerateOTPResponse, error)
	ResendOTP(otpReq ResendOTPReq) error
}
