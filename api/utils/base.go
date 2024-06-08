package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPaginationData(c *gin.Context) (int64, int64, error) {
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 0)

	if limit > 100 {
		limit = 100
	}
	if err != nil || limit < 0 {
		limit = 20
	}

	offset, err := strconv.ParseInt(c.Query("offset"), 10, 0)
	if err != nil || offset < 0 {
		offset = 0
	}

	return limit, offset, err
}
