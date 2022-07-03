package v1

import (
	"github.com/gin-gonic/gin"
)

type errResp struct {
	Error string `json:"error" example:"message"`
}

func newErrorResponse(c *gin.Context, code int, info any) {
	var msg string
	err, ok := info.(error)
	if ok {
		msg = err.Error()
	} else {
		msg = info.(string)
	}
	c.AbortWithStatusJSON(code, errResp{msg})
}
