package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/knightgao/dreame-api/common/helper"
)

func RequestId() func(c *gin.Context) {
	return func(c *gin.Context) {
		id := helper.GenRequestID()
		c.Set(helper.RequestIdKey, id)
		ctx := context.WithValue(c.Request.Context(), helper.RequestIdKey, id)
		c.Request = c.Request.WithContext(ctx)
		c.Header(helper.RequestIdKey, id)
		c.Next()
	}
}
