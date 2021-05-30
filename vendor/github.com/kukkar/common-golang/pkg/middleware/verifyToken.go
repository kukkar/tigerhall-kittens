package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/kukkar/common-golang/globalconst"
	"github.com/kukkar/common-golang/pkg/responsewriter"
	"github.com/kukkar/common-golang/pkg/sdk/merchant"
	"github.com/kukkar/common-golang/pkg/utils/clientvalidator"
	"github.com/kukkar/common-golang/pkg/utils/rError"
)

func (m *DefaultMiddleware) VerifyToken(clientValidator clientvalidator.ClientValidator) gin.HandlerFunc {

	return func(c *gin.Context) {

		var token, simID, installID, deviceID string
		if values, _ := c.Request.Header[globalconst.Token]; len(values) > 0 {
			token = values[0]
		}
		if values, _ := c.Request.Header[globalconst.HeaderDeviceID]; len(values) > 0 {
			deviceID = values[0]
		}
		if values, _ := c.Request.Header[globalconst.HeaderInstallID]; len(values) > 0 {
			installID = values[0]
		}
		if values, _ := c.Request.Header[globalconst.HeaderSimID]; len(values) > 0 {
			simID = values[0]
		}
		// sdk, err := merchant.GetMerchantSdk(merchantDBToUse)
		// if err != nil {
		// 	err = rError.MiscError(c, err, "Unable to get merchant SDK")
		// 	responsewriter.BuildResponse(c, "", err)
		// 	c.Abort()
		// 	return
		// }
		merchantInfo, err := clientValidator.VerifyToken(token)
		if err != nil {
			if err.Error() == merchant.UnAuthorisedUser.Error() {
				err = rError.UnauthoriseErr(c, err.Error())
				responsewriter.BuildResponse(c, "", err)
				c.Abort()
				return
			}
			err = rError.MiscError(c, err, "unable to verify token")
			responsewriter.BuildResponse(c, "", err)
			c.Abort()
			return
		}
		if simID != "" &&
			installID != "" &&
			deviceID != "" {
			err := clientValidator.VerifySimBinding(installID, simID, deviceID)
			if err != nil {
				err = rError.UnauthoriseErr(c, err.Error())
				responsewriter.BuildResponse(c, "", err)
				c.Abort()
				return
			}
		}
		c.Set(globalconst.MerchantInfo, merchantInfo)
		c.Next()
	}
}
