package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-app/internal/usecase"
	"todo-app/pkg/logger"
)

const (
	GetByIdURL = "/:uuid/"
	GetMeURL   = "/info/me/"
)

type accountHandler struct {
	useCase usecase.IAccountUseCase
	logger  logger.Interface
}

func newAccountRouter(router *gin.RouterGroup, useCase usecase.IAccountUseCase, log logger.Interface) {
	h := &accountHandler{
		useCase: useCase,
		logger:  log,
	}
	g := router.Group("/account")
	{
		g.GET(GetByIdURL, h.getById)
		g.GET(GetMeURL, h.getMe)
	}
}

func (h *accountHandler) getById(c *gin.Context) {
	_id := c.Param("uuid")
	account, err := h.useCase.GetById(c.Request.Context(), _id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, gin.H{
		"results": account,
	})
}

func (h *accountHandler) getMe(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	account, err := h.useCase.GetById(c.Request.Context(), userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, gin.H{
		"result": account,
	})
}
