package handlers

import "github.com/guatom999/backend-challenge/modules/middleware/usecases"

type (
	MiddlewareHandlerInterface interface {
	}

	middlewareHandler struct {
		middlewareUseCase usecases.MiddlewareUseCaseInterface
	}
)

func NewMiddlewareHandler(middlewareUseCase usecases.MiddlewareUseCaseInterface) MiddlewareHandlerInterface {
	return &middlewareHandler{middlewareUseCase: middlewareUseCase}
}
