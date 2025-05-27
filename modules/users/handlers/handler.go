package handlers

import (
	"context"
	"net/http"

	"github.com/guatom999/backend-challenge/modules/users"
	"github.com/guatom999/backend-challenge/modules/users/usecases"
	"github.com/labstack/echo/v4"
)

type (
	HandlerInterface interface {
		Register(c echo.Context) error
		GetAllUsers(c echo.Context) error
		CountUser(c echo.Context) error
		GetUserById(c echo.Context) error
		Login(c echo.Context) error
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

	req := new(users.CreateUserReq)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "error: invalid request body")
	}

	if err := h.usecase.Register(ctx, req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "success")
}

func (h *handler) Login(c echo.Context) error {

	ctx := context.Background()

	req := new(users.LoginCredentialReq)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "error: invalid request body")
	}

	userCredentail, err := h.usecase.Login(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, userCredentail)
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

	req := new(users.GetUserByIdReq)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "error: invalid request body")
	}

	user, err := h.usecase.GetUserById(ctx, req.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (h *handler) CountUser(c echo.Context) error {

	ctx := context.Background()

	result, err := h.usecase.CountUser(ctx)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)

}

func (h *handler) UpdateUserDetail(c echo.Context) error {

	ctx := context.Background()

	req := new(users.UpdateUserReq)

	userId := c.Get("user_id").(string)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "error: invalid request body")
	}

	if err := h.usecase.UpdateUserDetail(ctx, userId, req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "update success")

}

func (h *handler) DeleteUser(c echo.Context) error {

	ctx := context.Background()

	// req := new(users.UpdateUserReq)

	userId := c.Get("user_id").(string)

	// if err := c.Bind(req); err != nil {
	// 	return c.JSON(http.StatusBadRequest, "error: invalid request body")
	// }

	if err := h.usecase.Deleteuser(ctx, userId); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "delete success")

}
