package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/kukkar/common-golang/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kukkar/common-golang/globalconst"
	"github.com/kukkar/common-golang/pkg/logger"
	"github.com/kukkar/common-golang/pkg/responsewriter"
	"github.com/kukkar/common-golang/pkg/utils/rError"
)

func (m *DefaultMiddleware) CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, x-grant-code, x-access-token, Access-Token, access-token, RESPONSE-TYPE, response-type,token,deviceInfo,deviceinfo,DeviceInfo")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func (m *DefaultMiddleware) MonitorRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		setApiStartTime(c)

		uniqueKey, err := uniqueRequestKey(c, c.Request.Header, globalconst.ResponseHeaderUniqueID)
		if err != nil {
			err = rError.MiscError(c, err, "Unable to get unique id for api")
			responsewriter.BuildResponse(c, "", err)
			c.Abort()
			return
		}
		rc := utils.RequestContext{
			RequestID:    uniqueKey,
			Method:       c.Request.Method,
			URI:          c.FullPath(),
			APIStartTime: time.Now(),
			IP:           c.Request.RemoteAddr,
		}
		c.Set(globalconst.RequestContext, rc)
		logger.Logger.Info(fmt.Sprintf("[%s] MonitorRequest - [%s] -  %s   %v",
			uniqueKey, c.Request.Method, c.FullPath(), c.Request.Header))
		c.Next()
	}
}

func uniqueRequestKey(c *gin.Context, header http.Header, key string) (string, error) {
	var uniqueKey string
	if v := header.Get(key); v != "" {
		uniqueKey = v
	} else {
		uniqueID, err := uuid.NewRandom()
		if err != nil {
			return "", err
		}
		uniqueKey = strings.ReplaceAll(uniqueID.String(), "-", "")
	}
	c.Set(globalconst.UniqueAPIKey, uniqueKey)

	return uniqueKey, nil
}

func setApiStartTime(c *gin.Context) {
	c.Set(globalconst.ApiStartTime, time.Now())
}
