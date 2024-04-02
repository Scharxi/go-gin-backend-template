package routes

import (
	"github.com/BxfferOverflow/go-gin-template/controller"
	"github.com/gin-gonic/gin"
)

func AddSampleRoutes(group *gin.RouterGroup) {
	group.GET("/hello", controller.SayHello)
}
