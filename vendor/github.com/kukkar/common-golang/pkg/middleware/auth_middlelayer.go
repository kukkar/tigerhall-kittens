package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kukkar/common-golang/globalconst"
	"github.com/kukkar/common-golang/pkg/requestparser"
	"github.com/kukkar/common-golang/pkg/responsewriter"
	"github.com/kukkar/common-golang/pkg/utils/clientvalidator"
	"github.com/kukkar/common-golang/pkg/utils/rError"
)

func (m *DefaultMiddleware) ClientValidation(useClientValidator bool,
	cValidator clientvalidator.ClientValidator) gin.HandlerFunc {

	return func(c *gin.Context) {
		if useClientValidator {
			var clientName, hmac string
			if values, _ := c.Request.Header[globalconst.CLIENT_NAME]; len(values) > 0 {
				clientName = values[0]
			}
			if values, _ := c.Request.Header[globalconst.HMAC]; len(values) > 0 {
				hmac = values[0]
			}
			if clientName == "" || hmac == "" {
				err := rError.BadReqError(c, fmt.Errorf("client Name and hmac required"),
					"Unable to get client name and hmac in header")
				responsewriter.BuildResponse(c, "", err)
				c.Abort()
				return
			}
			data, err := buildData(c)
			if err != nil {
				err = rError.MiscError(c, err, "Unable to get request body")
				responsewriter.BuildResponse(c, "", err)
				c.Abort()
				return
			}
			validated, err := clientvalidator.ValidateHmac(data, hmac,
				clientName, cValidator)
			if err != nil {
				err = rError.MiscError(c, err)
				responsewriter.BuildResponse(c, "", err)
				c.Abort()
				return
			}
			if !validated {
				err = rError.UnauthoriseErr(c, "unauthorised client")
				responsewriter.BuildResponse(c, "", err)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

func buildData(c *gin.Context) (map[string]interface{}, error) {

	requestData := make(map[string]interface{})
	if c.Request.Method == "GET" {
		queryParam := c.Request.URL.Query()
		for key, value := range queryParam {
			if len(value) > 0 {
				requestData[key] = value[0]
			}
		}
	} else {
		err := requestparser.LoadBody(c, &requestData)
		if err != nil {
			return nil, err
		}
	}
	return requestData, nil
}
