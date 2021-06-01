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
	}
}
