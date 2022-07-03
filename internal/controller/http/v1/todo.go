package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-app/internal/entity"
	"todo-app/internal/usecase"
	"todo-app/pkg/logger"
)

const (
	TodoIdURL  = "/:uuid/"
	TodoGetAll = "/all"
)

type todoHandler struct {
	useCase usecase.ITodoUseCase
	logger  logger.Interface
}

func newTodoRouter(router *gin.RouterGroup, useCase usecase.ITodoUseCase, log logger.Interface) {
	h := &todoHandler{
		useCase: useCase,
		logger:  log,
	}
	g := router.Group("/todo")
	{
		g.POST("/", h.todoCreate)
		g.GET(TodoGetAll, h.todoGetAll)
		g.GET(TodoIdURL, h.todoGetById)
		g.PATCH(TodoIdURL, h.todoUpdate)
		g.DELETE(TodoIdURL, h.todoDelete)
	}
}

func (h *todoHandler) todoGetById(c *gin.Context) {
	_id := c.Param("uuid")
	res, err := h.useCase.GetById(c.Request.Context(), _id)
	if err != nil {
		newErrorResponse(c, http.StatusTeapot, err)
	}
	c.JSON(http.StatusOK, gin.H{"results": res})
}

func (h *todoHandler) todoGetAll(c *gin.Context) {
	accountId, _ := c.Get(userCtx)
	accountIdStr := accountId.(string)
	res, err := h.useCase.GetAll(c.Request.Context(), accountIdStr)
	if err != nil {
		newErrorResponse(c, http.StatusTeapot, err)
	}
	c.JSON(http.StatusOK, gin.H{"result": res})
}

func (h *todoHandler) todoCreate(c *gin.Context) {
	var dto *entity.TodoDTO
	if err := c.BindJSON(&dto); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("invalid input body, err: %v", err))
		return
	}
	accountId, _ := c.Get(userCtx)
	accountIdStr := accountId.(string)
	_id, err := h.useCase.Create(c.Request.Context(), dto, accountIdStr)
	if err != nil {
		newErrorResponse(c, http.StatusTeapot, fmt.Errorf("service errror. due to error: %v", err))
	}
	c.JSON(http.StatusOK, gin.H{"result": _id})
}

func (h *todoHandler) todoUpdate(c *gin.Context) {
	var dto *entity.TodoUpdateDTO
	if err := c.BindJSON(&dto); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("invalid input body, err: %v", err))
		return
	}
	_id := c.Param("uuid")
	res, err := h.useCase.Update(c.Request.Context(), dto, _id)
	if err != nil {
		newErrorResponse(c, http.StatusTeapot, fmt.Errorf("service error. due to error: %v", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": res})
}

func (h *todoHandler) todoDelete(c *gin.Context) {
	_id := c.Param("uuid")
	err := h.useCase.Delete(c.Request.Context(), _id)
	if err != nil {
		newErrorResponse(c, http.StatusTeapot, fmt.Errorf("service errror. due to error: %v", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}
