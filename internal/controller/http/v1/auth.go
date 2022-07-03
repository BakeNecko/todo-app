package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"todo-app/internal/entity"
	"todo-app/internal/usecase"
	"todo-app/pkg/logger"
)

const (
	SignInURL = "/sign-in/"
	SingUpURL = "/sign-up/"
)

type authHandler struct {
	useCase usecase.IAuthUseCase
	logger  logger.Interface
}

func newAuthRouter(router *gin.RouterGroup, useCase usecase.IAuthUseCase, log logger.Interface) {
	h := &authHandler{
		useCase: useCase,
		logger:  log,
	}
	g := router.Group("/auth")
	{
		g.POST(SignInURL, h.signIn)
		g.POST(SingUpURL, h.singUp)
	}
}

func (h *authHandler) signIn(c *gin.Context) {
	var dto entity.SignInDTO

	if err := c.BindJSON(&dto); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("invalid input body, err: %v", err))
		return
	}
	res, err := h.useCase.SignIn(c, &dto)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusTeapot, fmt.Sprintf("User doesn't exits"))
			return
		}
		h.logger.Error("failed to login User due to error: %v", err)
		newErrorResponse(c, http.StatusTeapot, fmt.Sprintf("Failed to login user. Error: %v", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": res})
}

func (h *authHandler) singUp(c *gin.Context) {
	var dto entity.AccountDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("invalid input body, err: %v", err))
		return
	}
	_id, err := h.useCase.SignUp(c.Request.Context(), &dto)
	if err != nil {
		h.logger.Error("failed to create User due to error: %v", err)
		newErrorResponse(c, http.StatusTeapot, err)
		return
	}
	c.JSON(201, gin.H{
		"id": _id,
	})
}
