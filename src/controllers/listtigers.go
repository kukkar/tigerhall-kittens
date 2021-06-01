package controllers

import (
	"fmt"

	"github.com/kukkar/common-golang/globalconst"

	appConf "github.com/kukkar/tigerhall-kittens/conf"
	"github.com/kukkar/tigerhall-kittens/src/tigerhall-kittens"

	"github.com/gin-gonic/gin"
	"github.com/kukkar/common-golang/pkg/responsewriter"
	"github.com/kukkar/common-golang/pkg/utils"
	"github.com/kukkar/common-golang/pkg/utils/rError"
)

//
//ListTigers service
func ListTigers(c *gin.Context) {

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

	tigerhallInstance, err := tigerhall.GetTigerHallKittens(c.Request.Context(), tigerhall.ConfigTigerHall{
		StorageAdapter: "mongo",
	})
	if err != nil {
		err = rError.MiscError(c, err, "Unable to get tiger hall instance")
		responsewriter.BuildResponse(c, "", err)
		return
	}
	queryList, err := parseQuery("")
	if err != nil {
		err = rError.MiscError(c, err, "Unable to get tiger hall instance")
		responsewriter.BuildResponse(c, "", err)
		return
	}
	data, count, err := tigerhallInstance.ListTigers(queryList, 10, 0)
	if err != nil {
		err = rError.MiscError(c, err, "Unable to get tiger hall instance")
		responsewriter.BuildResponse(c, "", err)
		return
	}
	fmt.Printf("## %v %v %v ", rc, conf.Mongo.DbName, gConfig)

	fmt.Printf("data %v count %v", data, count)
	responsewriter.BuildResponseWithBool(c, nil, nil)
	return
}
