package controllers

import (
	"fmt"
	"strconv"

	"github.com/kukkar/common-golang/globalconst"

	appConf "github.com/kukkar/tigerhall-kittens/conf"
	"github.com/kukkar/tigerhall-kittens/src/tigerhall-kittens"

	"github.com/gin-gonic/gin"
	"github.com/kukkar/common-golang/pkg/logger"
	"github.com/kukkar/common-golang/pkg/responsewriter"
	"github.com/kukkar/common-golang/pkg/utils"
	"github.com/kukkar/common-golang/pkg/utils/rError"
)

//
//ListTigerSight service
func ListTigerSight(c *gin.Context) {

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
	var limit, page int
	id := c.Query("id")
	limitString := c.Query("limit")
	pageString := c.Query("page")

	limit, err = strconv.Atoi(limitString)
	if err != nil {
		logger.Error(err, rc)
		limit = 10
	}
	page, err = strconv.Atoi(pageString)
	if err != nil {
		logger.Error(err, rc)
		page = 0
	}

	tigerhallInstance, err := tigerhall.GetTigerHallKittens(c.Request.Context(), tigerhall.ConfigTigerHall{
		StorageAdapter: "mongo",
	})
	if err != nil {
		err = rError.MiscError(c, err, "Unable to get tiger hall instance")
		responsewriter.BuildResponse(c, "", err)
		return
	}

	data, err := tigerhallInstance.ListSigntsOfTiger(id, limit, page)
	if err != nil {
		err = rError.MiscError(c, err, "Unable to get tiger hall instance")
		responsewriter.BuildResponse(c, "", err)
		return
	}
	fmt.Printf("## %v %v %v ", rc, conf.Mongo.DbName, gConfig)
	fmt.Printf("data %v ", data)
	responsewriter.BuildResponseWithBool(c, nil, nil)
	return
}
