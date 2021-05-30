package otp

func GetOTPSDK(config OTPConfig) (OTPSdk, error) {
	err := config.validateConfig()
	if err != nil {
		return nil, err
	}

	sdk := OTPAccessor{
		IPPort:          config.IPPort,
		Version:         config.Version,
		SuperKey:        config.SuperKey,
		MerchantDBToUse: config.MerchantDBToUse,
		clientvalidator: config.ClientValidator,
	}

	return sdk, nil
}
