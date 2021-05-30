package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/kukkar/common-golang/globalconst"
	"github.com/kukkar/common-golang/pkg/responsewriter"
	"github.com/kukkar/common-golang/pkg/sdk/otp"
)

func (m *DefaultMiddleware) OTPValidation(otpConfig otp.OTPConfig) gin.HandlerFunc {

	return func(c *gin.Context) {
		var clientName, mobileOTP, uuid string
		if values, _ := c.Request.Header[globalconst.CLIENT_NAME]; len(values) > 0 {
			clientName = values[0]
		}
		if values, _ := c.Request.Header[globalconst.OTPUUID]; len(values) > 0 {
			uuid = values[0]
		}
		if values, _ := c.Request.Header[globalconst.OTP]; len(values) > 0 {
			mobileOTP = values[0]
		}
		SDK, err := otp.GetOTPSDK(otpConfig)
		if err != nil {
			responsewriter.BuildResponse(c, "", err)
			c.Abort()
			return
		}
		err = SDK.VerfiyOTP(mobileOTP, uuid, clientName)
		if err != nil {
			responsewriter.BuildResponse(c, "", err)
			c.Abort()
			return
		}
		c.Next()
	}
}
