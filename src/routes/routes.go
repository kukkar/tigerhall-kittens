package routes

import (
	appConf "github.com/kukkar/tigerhall-kittens/conf"
	controller "github.com/kukkar/tigerhall-kittens/src/controllers"

	"github.com/gin-gonic/gin"
	"github.com/kukkar/common-golang/pkg/middleware"
)

func Routes(route *gin.Engine) {

	gConf, err := appConf.GetGlobalConfig()
	if err != nil {
		panic(err)
	}

	v1 := route.Group(string(gConf.AppName) + "/v1")
	{
		defaultMiddleware := middleware.DefaultMiddleware{}
		v1.GET("/listtigers", defaultMiddleware.MonitorRequest(), controller.ListTigers)
		v1.GET("/listtigersights", defaultMiddleware.MonitorRequest(), controller.ListTigerSight)
		v1.POST("/uploadimage", defaultMiddleware.MonitorRequest(), controller.UploadImage)
		v1.POST("/createtiger", defaultMiddleware.MonitorRequest(), controller.CreateTiger)
		v1.POST("/sighttiger", defaultMiddleware.MonitorRequest(), controller.CreateTigerSight)
	}
}
