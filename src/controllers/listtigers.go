package controllers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kukkar/common-golang/pkg/logger"
	"github.com/kukkar/common-golang/pkg/responsewriter"
	"github.com/kukkar/common-golang/pkg/utils/rError"
	"github.com/kukkar/tigerhall-kittens/src/tigerhall-kittens"
)

// ListTigers list tigers with provided query feature
// @Summary ListTigers list tigers with provided query feature
// @Produce json
// @Param q query string false "q"
// @Param limit query string false "limit"
// @Param page query string false "page"
// @Success 200 {object} ResListTigers
// @Router /v1/listtigers [get]
func ListTigers(c *gin.Context) {
	q := c.Query("q")

	tigerhallInstance, err := tigerhall.GetTigerHallKittens(c.Request.Context(), tigerhall.ConfigTigerHall{
		StorageAdapter: "mongo",
	})
	if err != nil {
		err = rError.MiscError(c, err, "Unable to get tiger hall instance")
		responsewriter.BuildResponse(c, "", err)
		return
	}
	var limit, page int
	limitString := c.Query("limit")
	pageString := c.Query("page")

	limit, err = strconv.Atoi(limitString)
	if err != nil {
		logger.Error(err)
		limit = 10
	}
	page, err = strconv.Atoi(pageString)
	if err != nil {
		logger.Error(err)
		page = 0
	}
	queryList, err := parseQuery(q)
	if err != nil {
		err = rError.MiscError(c, err, "Unable to get tiger hall instance")
		responsewriter.BuildResponse(c, "", err)
		return
	}
	data, count, err := tigerhallInstance.ListTigers(queryList, limit, page)
	if err != nil {
		err = rError.MiscError(c, err, "Unable to get tiger hall instance")
		responsewriter.BuildResponse(c, "", err)
		return
	}


	var serviceRes ResListTigers
	for _, eachData := range data {
		serviceRes.TigerData = append(serviceRes.TigerData, TigerData{
			ID:     eachData.ID,
			DOB:    eachData.DOB,
			SeenAt: eachData.SeenAt,
			Name:   eachData.Name,
		})
	}
	serviceRes.TotalCount = count
	fmt.Printf("data %v count %v", data, count)
	responsewriter.BuildResponseWithBool(c, serviceRes, nil)
	return
}
