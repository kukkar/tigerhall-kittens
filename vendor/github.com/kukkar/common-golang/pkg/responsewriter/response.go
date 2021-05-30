package responsewriter

import (
	"fmt"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/kukkar/common-golang/globalconst"
	"github.com/kukkar/common-golang/pkg/logger"
	"github.com/kukkar/common-golang/pkg/utils/rError"
	"go.elastic.co/apm"
)

//BuildResponse will build response
func BuildResponse(c *gin.Context, data interface{}, err error) {
	var apiKey string
	if uniqueID, ok := c.Get(globalconst.UniqueAPIKey); ok {
		apiKey = uniqueID.(string)
	} else {
		err = rError.MiscError(c, fmt.Errorf("unable to reterive api key"), "")
	}
	var rc interface{}
	var exist bool
	if rc, exist = c.Get(globalconst.RequestContext); !exist {
		logger.Error("unable to get request context " + err.Error())
	}
	resp := gin.H{}
	var statusCode HTTPCode
	statusCode = 200
	if err != nil {
		switch err.(type) {
		case *rError.Error:
			d, _ := err.(*rError.Error)
			statusCode = appErrorToHttpCodeMap[d.Code]
			resp["error"] = err
		default:
			resp["error"] = err.Error()
			statusCode = HTTPFatalErrorCode
		}
	} else {
		if data != nil {
			resp["data"] = data
		}
		resp["status"] = "success"
	}
	if statusCode == 206 {
		if data != nil {
			resp["data"] = data
		}
		resp["status"] = "partial"
	}
	if c.FullPath() != "/merchant/v2/listbank" &&
		c.FullPath() != "/merchant/v1/ifscfinder" {
		logger.Info(fmt.Sprintf("%+v ", resp), rc)
	}
	c.Header(globalconst.ResponseHeaderUniqueID, apiKey)
	c.JSON(int(statusCode), resp)
	if err != nil {
		pushInSentry(c, err)
		logger.Error(fmt.Sprintf("%s", err), rc)
	}
}
func BuildResponseMerchantINFO(c *gin.Context, data interface{}, err error) {
	var apiKey string
	if uniqueID, ok := c.Get(globalconst.UniqueAPIKey); ok {
		apiKey = uniqueID.(string)
	} else {
		err = rError.MiscError(c, fmt.Errorf("unable to reterive api key"), "")
	}
	var rc interface{}
	var exist bool
	if rc, exist = c.Get(globalconst.RequestContext); !exist {
		logger.Error("unable to get request context " + err.Error())
	}
	resp := gin.H{}
	var statusCode HTTPCode
	statusCode = 200
	if err != nil {
		switch err.(type) {
		case *rError.Error:
			d, _ := err.(*rError.Error)
			if appErrorToHttpCodeMap[d.Code] != 403 {
				statusCode = appErrorToHttpCodeMap[d.Code]
			}
			resp["error"] = err.Error()
			resp["alertmsg"] = err.Error()
		default:
			resp["error"] = err.Error()
			resp["alertmsg"] = err.Error()
			statusCode = HTTPFatalErrorCode
		}
		resp["response"] = "failed"
	} else {
		if data != nil {
			resp["data"] = data
		}
		resp["response"] = "success"
	}
	logger.Info(fmt.Sprintf("%+v", resp), rc)
	c.Header(globalconst.ResponseHeaderUniqueID, apiKey)
	c.JSON(int(statusCode), resp)
	if err != nil {
		pushInSentry(c, err)
		logger.Error(fmt.Sprintf("%s", err), rc)
	}
}

func BuildResponseOTPGenerator(c *gin.Context, data interface{}, err error) {
	var apiKey string
	if uniqueID, ok := c.Get(globalconst.UniqueAPIKey); ok {
		apiKey = uniqueID.(string)
	} else {
		err = rError.MiscError(c, fmt.Errorf("unable to reterive api key"), "")
	}
	var rc interface{}
	var exist bool
	if rc, exist = c.Get(globalconst.RequestContext); !exist {
		logger.Error("unable to get request context " + err.Error())
	}
	resp := gin.H{}
	var statusCode HTTPCode
	statusCode = 200
	if err != nil {
		switch err.(type) {
		case *rError.Error:
			d, _ := err.(*rError.Error)
			statusCode = appErrorToHttpCodeMap[d.Code]
			resp["error"] = d.Msg
		default:
			resp["error"] = err.Error()
			statusCode = HTTPFatalErrorCode
		}
	} else {
		resp["response"] = "Success"
	}
	logger.Info(fmt.Sprintf("%+v", resp), rc)
	c.Header(globalconst.ResponseHeaderUniqueID, apiKey)
	c.JSON(int(statusCode), resp)
	if err != nil {
		pushInSentry(c, err)
		logger.Error(fmt.Sprintf("%s", err), rc)
	}
}
func BuildResponseOTPVerify(c *gin.Context, data interface{}, err error) {
	var apiKey string
	if uniqueID, ok := c.Get(globalconst.UniqueAPIKey); ok {
		apiKey = uniqueID.(string)
	} else {
		err = rError.MiscError(c, fmt.Errorf("unable to reterive api key"), "")
	}
	var rc interface{}
	var exist bool
	if rc, exist = c.Get(globalconst.RequestContext); !exist {
		logger.Error("unable to get request context " + err.Error())
	}
	resp := gin.H{}
	var statusCode HTTPCode
	statusCode = 200
	if err != nil {
		switch err.(type) {
		case *rError.Error:
			d, _ := err.(*rError.Error)
			statusCode = appErrorToHttpCodeMap[d.Code]
			resp["error"] = d.Msg
		default:
			resp["error"] = err.Error()
			statusCode = HTTPFatalErrorCode
		}
		logger.Logger.Info(fmt.Sprintf("%s Api Response %+v", apiKey, resp))
		c.JSON(int(statusCode), resp)
		return
	}
	logger.Info(fmt.Sprintf("%+v ", resp), rc)
	c.Header(globalconst.ResponseHeaderUniqueID, apiKey)
	c.JSON(int(statusCode), data)
	if err != nil {
		pushInSentry(c, err)
		logger.Error(fmt.Sprintf("%s", err), rc)
	}
}

func BuildResponseLogout(c *gin.Context, data interface{}, err error) {
	var apiKey string
	if uniqueID, ok := c.Get(globalconst.UniqueAPIKey); ok {
		apiKey = uniqueID.(string)
	} else {
		err = rError.MiscError(c, fmt.Errorf("unable to reterive api key"), "")
	}
	var rc interface{}
	var exist bool
	if rc, exist = c.Get(globalconst.RequestContext); !exist {
		logger.Error("unable to get request context " + err.Error())
	}
	resp := gin.H{}
	var statusCode HTTPCode
	statusCode = 200
	if err != nil {
		switch err.(type) {
		case *rError.Error:
			d, _ := err.(*rError.Error)
			statusCode = appErrorToHttpCodeMap[d.Code]
			resp["error"] = d.Msg
		default:
			resp["error"] = err.Error()
			statusCode = HTTPFatalErrorCode
		}
		logger.Logger.Error(fmt.Sprintf("%s Api Response %+v", apiKey, resp))
		c.JSON(int(statusCode), resp)
		return
	}
	logger.Info(fmt.Sprintf("%+v ", resp), rc)
	c.Header(globalconst.ResponseHeaderUniqueID, apiKey)
	c.JSON(int(statusCode), data)
	if err != nil {
		pushInSentry(c, err)
		logger.Error(fmt.Sprintf("%s", err), rc)
	}
}

func BuildResponseWithBool(c *gin.Context, data interface{}, err error) {
	var apiKey string
	if uniqueID, ok := c.Get(globalconst.UniqueAPIKey); ok {
		apiKey = uniqueID.(string)
	} else {
		err = rError.MiscError(c, fmt.Errorf("unable to reterive api key"), "")
	}
	var rc interface{}
	var exist bool
	if rc, exist = c.Get(globalconst.RequestContext); !exist {
		logger.Error("unable to get request context " + err.Error())
	}
	resp := gin.H{}
	var statusCode HTTPCode
	statusCode = 200
	if err != nil {
		switch err.(type) {
		case *rError.Error:
			d, _ := err.(*rError.Error)
			statusCode = appErrorToHttpCodeMap[d.Code]
			resp["error"] = err
			resp["status"] = false
		default:
			resp["error"] = err.Error()
			statusCode = HTTPFatalErrorCode
			resp["status"] = false
		}
	} else {
		if data != nil {
			resp["data"] = data
		}
		resp["status"] = true
	}
	if statusCode == 206 {
		if data != nil {
			resp["data"] = data
		}
		resp["status"] = true
	}
	logger.Info(fmt.Sprintf("%+v ", resp), rc)
	c.Header(globalconst.ResponseHeaderUniqueID, apiKey)
	c.JSON(int(statusCode), resp)
	if err != nil {
		pushInSentry(c, err)
		logger.Error(fmt.Sprintf("%s", err), rc)
	}
}
func BuildResponseDeviceAlreadyRegister(c *gin.Context, data interface{}, err error) {
	var apiKey string
	if uniqueID, ok := c.Get(globalconst.UniqueAPIKey); ok {
		apiKey = uniqueID.(string)
	} else {
		err = rError.MiscError(c, fmt.Errorf("unable to reterive api key"), "")
	}
	var rc interface{}
	var exist bool
	if rc, exist = c.Get(globalconst.RequestContext); !exist {
		logger.Error("unable to get request context " + err.Error())
	}
	resp := gin.H{}
	var statusCode HTTPCode
	statusCode = 200
	if err != nil {
		switch err.(type) {
		case *rError.Error:
			d, _ := err.(*rError.Error)
			statusCode = appErrorToHttpCodeMap[d.Code]
			resp["alertmsg"] = "You are not able to register.Merchant Already Registered"
			resp["response"] = "failed"
		default:
			resp["error"] = err.Error()
			statusCode = HTTPFatalErrorCode
		}
	} else {
		if data != nil {
			resp["data"] = data
		}
		resp["status"] = "success"
	}

	logger.Info(fmt.Sprintf("%+v ", resp), rc)
	c.Header(globalconst.ResponseHeaderUniqueID, apiKey)
	c.JSON(int(statusCode), resp)
	if err != nil {
		pushInSentry(c, err)
		logger.Error(fmt.Sprintf("%s", err), rc)
	}
}

func pushInSentry(c *gin.Context, err error) {
	var apiKey string
	if uniqueID, ok := c.Get(globalconst.UniqueAPIKey); ok {
		apiKey = uniqueID.(string)
	}
	var rErrorRes rError.Error
	switch err.(type) {
	case *rError.Error:
		errData, typeErr := err.(*rError.Error)
		if !typeErr {
			rErrorRes.Msg = err.Error()
		}
		rErrorRes = (*errData)
	default:
		rErrorRes.Msg = err.Error()
	}
	rErrorRes.Info["UniqueID"] = apiKey
	apm.CaptureError(c.Request.Context(), err).Send()
	sentry.CaptureException(&rErrorRes)
}
func BuildResponsForTransaction(c *gin.Context, data interface{}, err error) {
	var apiKey string
	if uniqueID, ok := c.Get(globalconst.UniqueAPIKey); ok {
		apiKey = uniqueID.(string)
	} else {
		err = rError.MiscError(c, fmt.Errorf("unable to reterive api key"), "")
	}
	var rc interface{}
	var exist bool
	if rc, exist = c.Get(globalconst.RequestContext); !exist {
		logger.Error("unable to get request context " + err.Error())
	}
	resp := gin.H{}
	var statusCode HTTPCode
	statusCode = 200
	if err != nil {
		switch err.(type) {
		case *rError.Error:
			d, _ := err.(*rError.Error)
			statusCode = appErrorToHttpCodeMap[d.Code]
			resp["msg"] = "No UPI Txns. in last 15 days"
			resp["success"] = false
		default:
			resp["error"] = err.Error()
			statusCode = HTTPFatalErrorCode
			resp["success"] = false
		}
	} else {
		if data != nil {
			resp["txns_details"] = data
		}
		resp["success"] = true
	}
	if statusCode == 206 {
		if data != nil {
			resp["data"] = data
		}
		resp["success"] = true
	}
	logger.Info(fmt.Sprintf("%+v ", resp), rc)
	c.Header(globalconst.ResponseHeaderUniqueID, apiKey)
	c.JSON(int(statusCode), resp)
	if err != nil {
		pushInSentry(c, err)
		logger.Error(fmt.Sprintf("%s", err), rc)
	}
}
func BuildResponsForSettlementV1(c *gin.Context, data interface{}, amount *float32, err error) {
	var apiKey string
	if uniqueID, ok := c.Get(globalconst.UniqueAPIKey); ok {
		apiKey = uniqueID.(string)
	} else {
		err = rError.MiscError(c, fmt.Errorf("unable to reterive api key"), "")
	}
	var rc interface{}
	var exist bool
	if rc, exist = c.Get(globalconst.RequestContext); !exist {
		logger.Error("unable to get request context " + err.Error())
	}
	resp := gin.H{}
	var statusCode HTTPCode
	statusCode = 200
	if err != nil {
		switch err.(type) {
		case *rError.Error:
			d, _ := err.(*rError.Error)
			statusCode = appErrorToHttpCodeMap[d.Code]
			resp["msg"] = "No Settlement in Last 15 Days"
			resp["success"] = false
		default:
			resp["error"] = err.Error()
			statusCode = HTTPFatalErrorCode
			resp["success"] = false
		}
	} else {
		if data != nil {
			resp["settlement_details"] = data
			resp["total_amount"] = amount
		}
		resp["success"] = true
	}
	if statusCode == 206 {
		if data != nil {
			resp["data"] = data
		}
		resp["success"] = true
	}
	logger.Info(fmt.Sprintf("%+v ", resp), rc)
	c.Header(globalconst.ResponseHeaderUniqueID, apiKey)
	c.JSON(int(statusCode), resp)
	if err != nil {
		pushInSentry(c, err)
		logger.Error(fmt.Sprintf("%s", err), rc)
	}
}
