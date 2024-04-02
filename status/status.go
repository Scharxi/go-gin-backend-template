package status

import (
	"github.com/BxfferOverflow/gogintemplate/httperror"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ok(c *gin.Context, value any) {
	c.JSON(http.StatusOK, value)
}

func Created(c *gin.Context, value any) {
	c.JSON(http.StatusCreated, value)
}

func Unauthorized(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, httperror.NewHttpError(err))
}

func NotFound(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusNotFound, httperror.NewHttpError(err))
}

func Error(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, httperror.NewHttpError(err))
}

func BadRequest(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusBadRequest, httperror.NewHttpError(err))
}
