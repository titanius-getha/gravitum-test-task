package userhandler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/titanius-getha/gravitum-test-task/domain/user"
	transport "github.com/titanius-getha/gravitum-test-task/pkg/transport/http"
)

type UserProvider interface {
	GetByID(ID uint) (user.User, error)
	Create(name string) (user.User, error)
	Update(ID uint, name string) (user.User, error)
}

type UserHandler struct {
	provider UserProvider
}

func New(provider UserProvider) *UserHandler {
	return &UserHandler{provider}
}

func (h *UserHandler) GetUser(ctx *gin.Context) {
	strID := ctx.Param("id")

	ID, err := strconv.Atoi(strID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, transport.BadResponse(transport.ErrBadRequest.Error()))
		return
	}

	u, err := h.provider.GetByID(uint(ID))
	if errors.Is(err, user.ErrUserNotFound) {
		ctx.JSON(http.StatusNotFound, transport.BadResponse(err.Error()))
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, transport.BadResponse(transport.ErrInternal.Error()))
		return
	}

	ctx.JSON(http.StatusOK, transport.GoodResponse(u))
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var body CreateUserDto
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, transport.BadResponse(transport.ErrBadRequest.Error()))
		return
	}

	u, err := h.provider.Create(body.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, transport.BadResponse(transport.ErrInternal.Error()))
		return
	}

	ctx.JSON(http.StatusOK, transport.GoodResponse(u))
}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	strID := ctx.Param("id")

	ID, err := strconv.Atoi(strID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, transport.BadResponse(transport.ErrBadRequest.Error()))
		return
	}

	var body UpdateUserDto
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, transport.BadResponse(transport.ErrBadRequest.Error()))
		return
	}

	u, err := h.provider.Update(uint(ID), body.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, transport.BadResponse(transport.ErrInternal.Error()))
		return
	}

	ctx.JSON(http.StatusOK, transport.GoodResponse(u))
}
