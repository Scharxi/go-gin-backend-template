package httputil

import (
	"context"
	"github.com/BxfferOverflow/gogintemplate/httperror"
	"github.com/BxfferOverflow/gogintemplate/status"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func Id(c *gin.Context, key string) int32 {
	value := c.Param(key)

	id, err := strconv.Atoi(value)
	if err != nil {
		status.BadRequest(c, httperror.InvalidIdType)
	}

	return int32(id)
}

func FromParam[T interface{}](c *gin.Context, key string, valid func(context.Context, string) (T, error)) *T {
	value, err := valid(c, c.Param(key))
	go log.Printf("%v", err)
	if err != nil {
		status.BadRequest(c, err)
		return nil
	}

	return &value
}
