package usecases

import (
	"github.com/guatom999/backend-challenge/modules/middleware/repositories"
	"github.com/labstack/echo/v4"
)

type (
	MiddlewareUseCaseInterface interface {
	}

	middlewareUsecase struct {
		middlewareRepo repositories.MiddlewareRepositoryInterface
	}
)

func NewMiddlewareUsecase(middlewareRepo repositories.MiddlewareRepositoryInterface) MiddlewareUseCaseInterface {
	return &middlewareUsecase{middlewareRepo: middlewareRepo}
}

func (u *middlewareUsecase) JwtAuthen(c echo.Context) (echo.Context, error) {

	// ctx := c.Request().Context()

	return nil, nil

}
