package v1

import (
	"github.com/gin-gonic/gin"
	"todo-app/internal/usecase"
	"todo-app/pkg/logger"
)

type Deps struct {
	UseCases        *usecase.UseCases
	JwtTokenManager usecase.ITokenManager
}

func NewRouter(handler *gin.Engine, l logger.Interface, deps *Deps) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Routers
	v1 := handler.Group("/api")
	v1AuthRequired := v1.Group("/v1", deps.authMiddleware)
	{
		newAuthRouter(v1, deps.UseCases.AuthUseCase, l)
		newAccountRouter(v1AuthRequired, deps.UseCases.AccountUseCase, l)
		newTodoRouter(v1AuthRequired, deps.UseCases.TodoUseCase, l)
	}
}
