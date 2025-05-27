package middlewareHandlers

import (
	"net/http"
	"strings"

	"github.com/guatom999/backend-challenge/modules/middleware/middlewareUsecases"
	"github.com/labstack/echo/v4"
)

type (
	MiddlewareHandlerInterface interface {
		JwtAuthentication(next echo.HandlerFunc) echo.HandlerFunc
	}

	middlewareHandler struct {
		middlewareUseCase middlewareUsecases.MiddlewareUseCaseInterface
	}
)

func NewMiddlewareHandler(middlewareUseCase middlewareUsecases.MiddlewareUseCaseInterface) MiddlewareHandlerInterface {
	return &middlewareHandler{middlewareUseCase: middlewareUseCase}
}

func (h *middlewareHandler) JwtAuthentication(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		token := strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer ")

		c, err := h.middlewareUseCase.JwtAuthen(c, token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		return next(c)
	}

}
