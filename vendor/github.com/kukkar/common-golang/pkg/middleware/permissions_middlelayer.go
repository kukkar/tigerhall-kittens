package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/kukkar/common-golang/globalconst"
	"github.com/kukkar/common-golang/pkg/responsewriter"
	"github.com/kukkar/common-golang/pkg/utils/clientvalidator"
	"github.com/kukkar/common-golang/pkg/utils/rError"
)

func (m *DefaultMiddleware) PermissionValidation(clientValidator clientvalidator.ClientValidator) gin.HandlerFunc {

	return func(c *gin.Context) {

		var clientName string
		if values, _ := c.Request.Header[globalconst.CLIENT_NAME]; len(values) > 0 {
			clientName = values[0]
		}

		// merchantSDK, err := merchant.GetMerchantSdk(merchantDBToUse)
		// if err != nil {
		// 	err = rError.MiscError(c, err)
		// 	responsewriter.BuildResponse(c, "", err)
		// 	c.Abort()
		// 	return
		// }
		permissions, err := clientValidator.GetClientPermissions(clientName)
		if err != nil {
			err = rError.MiscError(c, err)
			responsewriter.BuildResponse(c, "", err)
			c.Abort()
			return
		}
		if _, ok := permissions[c.FullPath()]; !ok {
			err = rError.UnauthoriseErr(c, "You don't have permission to use this route contact administrator")
			responsewriter.BuildResponse(c, "", err)
			c.Abort()
			return
		}
		c.Set(globalconst.ContKeyInternalClientPermissions, permissions)
		c.Next()
	}
}
