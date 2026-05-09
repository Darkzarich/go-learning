package handler

import (
	"net/http"
	"strconv"

	"users-service/internal/service"
	"users-service/pkg/apperror"
	"users-service/pkg/httputil"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

// RegisterRoutes sets up the HTTP routes on a ServeMux.
func (h *UserHandler) RegisterRoutes(g *gin.RouterGroup) {
	users := g.Group("/users")
	users.POST("", h.handleCreateUser)
	users.GET("", h.handleListUsers)
	users.GET("/:id", h.handleUserByID)
	users.PUT("/:id", h.handleUpdateUser)
	users.DELETE("/:id", h.handleDeleteUser)
}

func (h *UserHandler) handleCreateUser(c *gin.Context) {
	var CreateUserInput struct {
		Name  string `json:"name" binding:"required,min=3,max=20"`
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&CreateUserInput); err != nil {
		httputil.RespondValidationError(c, err)
		return
	}

	user, err := h.svc.Create(CreateUserInput.Name, CreateUserInput.Email)
	if err != nil {
		httputil.RespondError(c, err)
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) handleListUsers(c *gin.Context) {
	users, err := h.svc.GetAll()
	if err != nil {
		httputil.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) handleUserByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		httputil.RespondError(c, apperror.NewInvalidInput("invalid user id"))
		return
	}

	user, err := h.svc.GetByID(id)
	if err != nil {
		httputil.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) handleUpdateUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		httputil.RespondError(c, apperror.NewInvalidInput("invalid user id"))
		return
	}

	var UpdateUserInput struct {
		Name  string `json:"name" binding:"required,min=3,max=20"`
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&UpdateUserInput); err != nil {
		httputil.RespondValidationError(c, err)
		return
	}

	user, err := h.svc.Update(id, UpdateUserInput.Name, UpdateUserInput.Email)
	if err != nil {
		httputil.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) handleDeleteUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		httputil.RespondError(c, apperror.NewInvalidInput("invalid user id"))
		return
	}

	err = h.svc.DeleteByID(id)
	if err != nil {
		httputil.RespondError(c, err)
		return
	}

	httputil.RespondOK(c)
}
