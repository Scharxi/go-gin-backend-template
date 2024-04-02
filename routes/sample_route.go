package routes

import (
	"github.com/BxfferOverflow/gogintemplate/handler"
	"github.com/gin-gonic/gin"
)

func AddSampleRoutes(group *gin.RouterGroup) {
	group.GET("/hello", handler.SayHello)
}
