package controllers

import (
	"fmt"
	"time"

	appConf "github.com/kukkar/tigerhall-kittens/conf"
	"github.com/kukkar/tigerhall-kittens/src/tigerhall-kittens"

	"github.com/gin-gonic/gin"
	"github.com/kukkar/common-golang/globalconst"
	"github.com/kukkar/common-golang/pkg/requestparser"
	"github.com/kukkar/common-golang/pkg/responsewriter"
	"github.com/kukkar/common-golang/pkg/utils"
	"github.com/kukkar/common-golang/pkg/utils/rError"
)

// CreateTigerSight create sight of a tiger in the wild
// @Summary Createtiger create tiger in the wild
// @Produce json
// @Param add requesttigersight body ReqSightATiger true "create sight of tiger in wild"
// @Success 200
// @Router /v1/sighttiger [post]
func CreateTigerSight(c *gin.Context) {

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

	var req ReqSightATiger
	err = requestparser.LoadBody(c, &req)
	if err != nil {
		err = rError.UnmarshalError(c, err, "unable to unmarshal request")
		responsewriter.BuildResponseWithBool(c, "", err)
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

	err = tigerhallInstance.SightATiger(req.TigerID, tigerhall.ReqSightOfATiger{
		ImagePath: req.ImagePath,
		Coordinates: tigerhall.Coordinates{
			Lat:  req.Coordinates.Lat,
			Long: req.Coordinates.Long,
		},
		TimeStamp: time.Now(),
	})
	if err != nil {
		err = rError.MiscError(c, err, "Unable to create sight of tiger")
		responsewriter.BuildResponse(c, "", err)
		return
	}

	fmt.Printf("## %v %v %v ", rc, conf.Mongo.DbName, gConfig)
	responsewriter.BuildResponseWithBool(c, nil, nil)
	return
}
