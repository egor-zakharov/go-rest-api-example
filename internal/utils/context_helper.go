package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetInt64FromContext(c *gin.Context, param string) (int64, error) {
	paramId := c.Param(param)
	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}
