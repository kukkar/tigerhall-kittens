package controllers

import (
	"fmt"

	"github.com/kukkar/common-golang/globalconst"

	appConf "github.com/kukkar/tigerhall-kittens/conf"

	"github.com/gin-gonic/gin"
	"github.com/kukkar/common-golang/pkg/responsewriter"
	"github.com/kukkar/common-golang/pkg/utils"
	"github.com/kukkar/common-golang/pkg/utils/rError"
)

//
//HelloWorld service
func HelloWorld(c *gin.Context) {

	var rc utils.RequestContext
	if requestContext, ok := c.Get(globalconst.RequestContext); ok {
		rc = requestContext.(utils.RequestContext)
	}
	conf, err := appConf.GetAppConfig()
	if err != nil {
		err = rError.MiscError(c, err, "Unable to get appconfig")
		responsewriter.BuildResponse(c, "", err)
		return
	}
	gConfig, err := appConf.GetGlobalConfig()
	if err != nil {
		err = rError.MiscError(c, err, "Unable to get appconfig")
		responsewriter.BuildResponse(c, "", err)
		return
	}
	fmt.Printf("## %v %v %v ", rc, conf, gConfig)

	responsewriter.BuildResponseWithBool(c, nil, nil)
	return
}
