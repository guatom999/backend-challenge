package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/guatom999/backend-challenge/config"
	"github.com/guatom999/backend-challenge/modules/middleware/middlewareHandlers"
	"github.com/guatom999/backend-challenge/modules/middleware/middlewareRepositories"
	"github.com/guatom999/backend-challenge/modules/middleware/middlewareUsecases"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	Server interface {
		Start(pctx context.Context)
	}

	server struct {
		app        *echo.Echo
		cfg        *config.Config
		db         *mongo.Client
		middleware middlewareHandlers.MiddlewareHandlerInterface
	}
)

func NewMiddleware(cfg *config.Config, db *mongo.Client) middlewareHandlers.MiddlewareHandlerInterface {
	middlewareRepo := middlewareRepositories.NewRepository(db)
	middlewareUseCase := middlewareUsecases.NewMiddlewareUsecase(cfg, middlewareRepo)
	middlewareHandler := middlewareHandlers.NewMiddlewareHandler(middlewareUseCase)

	return middlewareHandler

}

func NewEchoServer(db *mongo.Client, cfg *config.Config) Server {
	return &server{
		app:        echo.New(),
		cfg:        cfg,
		db:         db,
		middleware: NewMiddleware(cfg, db),
	}
}

func (s *server) gratefulShutdown(pctx context.Context, close <-chan os.Signal) {
	<-close

	if err := s.app.Shutdown(pctx); err != nil {
		log.Fatal("Failed to shutdown Server....")
	}

	log.Println("Shutting Down Server....")
}

func (s *server) Start(pctx context.Context) {

	s.app.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "Error: Request Timeout",
		Timeout:      time.Second * 10,
	}))

	s.app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	s.app.Use(middleware.Logger())

	close := make(chan os.Signal, 1)
	signal.Notify(close, syscall.SIGINT, syscall.SIGTERM)

	go s.gratefulShutdown(pctx, close)

	s.UserService()

	if err := s.app.Start(s.cfg.App.Port); err != nil && err != http.ErrServerClosed {
		log.Printf("Failed to start Server %v ", err)
	}

}
