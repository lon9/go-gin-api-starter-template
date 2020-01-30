package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func parseLimitOffset(c *gin.Context) (offset, limit int, err error) {
	offsetStr := c.DefaultQuery("offset", "0")
	offset, err = strconv.Atoi(offsetStr)
	if err != nil {
		return
	}
	limitStr := c.DefaultQuery("limit", "5")
	limit, err = strconv.Atoi(limitStr)
	if limit > 50 {
		limit = 50
	}
	return
}
