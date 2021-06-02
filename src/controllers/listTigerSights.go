package controllers

import (
	"strconv"

	"github.com/kukkar/common-golang/globalconst"

	"github.com/kukkar/tigerhall-kittens/src/tigerhall-kittens"

	"github.com/gin-gonic/gin"
	"github.com/kukkar/common-golang/pkg/logger"
	"github.com/kukkar/common-golang/pkg/responsewriter"
	"github.com/kukkar/common-golang/pkg/utils"
	"github.com/kukkar/common-golang/pkg/utils/rError"
)

// ListTigerSight list tigers previous sights
// @Summary ListTigerSight list tigers previous sights
// @Produce json
// @Param limit query string false "limit"
// @Param page query string false "page"
// @Param id query string true "id"
// @Success 200 {object} ResListTigers
// @Router /v1/listtigersights [get]
func ListTigerSight(c *gin.Context) {

	var rc utils.RequestContext
	if requestContext, ok := c.Get(globalconst.RequestContext); ok {
		rc = requestContext.(utils.RequestContext)
	}

	var limit, page int
	id := c.Query("id")
	limitString := c.Query("limit")
	pageString := c.Query("page")

	limit, err := strconv.Atoi(limitString)
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
	var res ResSightATiger

	data, count, err := tigerhallInstance.ListSigntsOfTiger(id, limit, page)
	if err != nil {
		err = rError.MiscError(c, err, "Unable to get tiger hall instance")
		responsewriter.BuildResponse(c, "", err)
		return
	}
	res.TotalCount = *count
	for _, eachData := range data.TigerSights {
		res.Data = append(res.Data, ResSightData{
			Coordinates: Coordinates{
				Lat:  eachData.Coordinates.Lat,
				Long: eachData.Coordinates.Long,
			},
			ImagePath: eachData.ImagePath,
			SeenAt:    eachData.TimeStamp,
		})
	}
	responsewriter.BuildResponseWithBool(c, res, nil)
	return
}
