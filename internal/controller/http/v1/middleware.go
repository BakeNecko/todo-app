package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (d *Deps) authMiddleware(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, errors.New("empty auth header"))
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, errors.New("invalid auth header"))
		return
	}

	if len(headerParts[1]) == 0 {
		newErrorResponse(c, http.StatusUnauthorized, errors.New("token is empty"))
		return
	}

	userID, err := d.JwtTokenManager.Parse(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err)
		return
	}

	c.Set(userCtx, userID)
}

func getUserId(c *gin.Context) (string, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return "", errors.New("user id not found")
	}

	idStr, ok := id.(string)
	if !ok {
		return "", errors.New("user id is of invalid type")
	}

	return idStr, nil
}
