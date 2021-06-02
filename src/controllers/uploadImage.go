package controllers

import (
	appConf "github.com/kukkar/tigerhall-kittens/conf"
	"github.com/kukkar/tigerhall-kittens/src/tigerhall-kittens"

	"github.com/gin-gonic/gin"
	"github.com/kukkar/common-golang/pkg/requestparser"
	"github.com/kukkar/common-golang/pkg/responsewriter"
	"github.com/kukkar/common-golang/pkg/utils/rError"
)

// UploadImage upload image add images in variations depending on requirement
// @Summary upload image add images in variations depending on requirement
// @Produce json
// @Param add uploadimage body ReqUploadImage true "upload a image"
// @Success 200 {object} ResUploadImage
// @Router /v1/uploadimage [post]
func UploadImage(c *gin.Context) {

	conf, err := appConf.GetAppConfig()
	if err != nil {
		err = rError.MiscError(c, err, "Unable to get appconfig")
		responsewriter.BuildResponse(c, "", err)
		return
	}

	var req ReqUploadImage
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

	image, err := req.toImage(conf)
	if err != nil {
		err = rError.MiscError(c, err, "Unable to convert image data into image struct")
		responsewriter.BuildResponse(c, "", err)
		return
	}
	imagePath, err := tigerhallInstance.CreateImage(image)
	if err != nil {
		err = rError.MiscError(c, err, "unable to upload image")
		responsewriter.BuildResponse(c, "", err)
		return
	}
	if imagePath == nil {
		err = rError.MiscError(c, err, "Unable to get image path from service")
		responsewriter.BuildResponse(c, "", err)
		return
	}

	responsewriter.BuildResponseWithBool(c, ResUploadImage{
		ImagePath: *imagePath,
	}, nil)
	return
}
