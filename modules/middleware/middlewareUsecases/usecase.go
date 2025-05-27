package middlewareUsecases

import (
	"log"

	"github.com/guatom999/backend-challenge/config"
	"github.com/guatom999/backend-challenge/modules/middleware/middlewareRepositories"
	"github.com/guatom999/backend-challenge/pkg/jwtauth"
	"github.com/labstack/echo/v4"
)

type (
	MiddlewareUseCaseInterface interface {
		JwtAuthen(c echo.Context, token string) (echo.Context, error)
	}

	middlewareUsecase struct {
		cfg            *config.Config
		middlewareRepo middlewareRepositories.MiddlewareRepositoryInterface
	}
)

func NewMiddlewareUsecase(cfg *config.Config, middlewareRepo middlewareRepositories.MiddlewareRepositoryInterface) MiddlewareUseCaseInterface {
	return &middlewareUsecase{cfg: cfg, middlewareRepo: middlewareRepo}
}

func (u *middlewareUsecase) JwtAuthen(c echo.Context, token string) (echo.Context, error) {

	ctx := c.Request().Context()

	claims, err := jwtauth.ParseToken(ctx, u.cfg.Jwt.Secret, token)
	if err != nil {
		log.Printf("Erorr is %s", err.Error())
		return c, err
	}

	c.Set("user_id", claims.Claims.UserId)

	return c, nil

}
