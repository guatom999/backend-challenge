package handlers

import (
	"context"
	"net/http"

	"github.com/guatom999/backend-challenge/modules"
	"github.com/guatom999/backend-challenge/modules/usecases"
	"github.com/labstack/echo/v4"
)

type (
	HandlerInterface interface {
		Register(c echo.Context) error
		GetAllUsers(c echo.Context) error
		GetUserById(c echo.Context) error
		UpdateUserDetail(c echo.Context) error
		DeleteUser(c echo.Context) error
	}

	handler struct {
		usecase usecases.UsecaseInterface
	}
)

func NewHandler(usecase usecases.UsecaseInterface) HandlerInterface {
	return &handler{usecase: usecase}
}

func (h *handler) Register(c echo.Context) error {

	ctx := context.Background()

	req := new(modules.CreateUserReq)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "error: invalid request body")
	}

	if err := h.usecase.Register(ctx, req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "success")
}

func (h *handler) GetAllUsers(c echo.Context) error {

	ctx := context.Background()

	users, err := h.usecase.GetAllUses(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func (h *handler) GetUserById(c echo.Context) error {

	ctx := context.Background()

	req := new(modules.GetUserByIdReq)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "error: invalid request body")
	}

	user, err := h.usecase.GetUserById(ctx, req.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (h *handler) UpdateUserDetail(c echo.Context) error {

	ctx := context.Background()

	req := new(modules.UpdateUserReq)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "error: invalid request body")
	}

	if err := h.usecase.UpdateUserDetail(ctx, req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return nil

}

func (h *handler) DeleteUser(c echo.Context) error {

	ctx := context.Background()

	req := new(modules.UpdateUserReq)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "error: invalid request body")
	}

	if err := h.usecase.UpdateUserDetail(ctx, req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return nil

}
