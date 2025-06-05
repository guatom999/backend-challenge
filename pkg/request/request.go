package request

import (
	"errors"
	"log"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type (
	contextWrapperInterface interface {
		Bind(data any) error
	}

	contextWrapper struct {
		Context   echo.Context
		validator validator.Validate
	}
)

func NewContextWrapper(
	Context echo.Context,
	validator validator.Validate,
) contextWrapperInterface {
	return &contextWrapper{
		Context:   Context,
		validator: validator,
	}
}

func (cw *contextWrapper) Bind(data any) error {

	if err := cw.Context.Bind(data); err != nil {
		log.Printf("Error: Bind Data failed: %s", err.Error())
		return errors.New("error: bad request")
	}

	if err := cw.validator.Struct(data); err != nil {
		log.Printf("Error: Validate Struct failed: %s", err.Error())
		return errors.New("error: bad request")
	}

	return nil
}
