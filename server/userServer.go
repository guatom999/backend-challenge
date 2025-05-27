package server

import (
	"context"
	"log"
	"time"

	"github.com/guatom999/backend-challenge/modules/users/handlers"
	"github.com/guatom999/backend-challenge/modules/users/repositories"
	"github.com/guatom999/backend-challenge/modules/users/usecases"
)

func (s *server) UserService() {

	userRepository := repositories.NewRepository(s.db)
	userUseCase := usecases.NewUseCase(s.cfg, userRepository)
	userHanlder := handlers.NewHandler(userUseCase)

	route := s.app.Group("/user")

	route.POST("/register", userHanlder.Register)
	route.POST("/login", userHanlder.Login)

	route.GET("/listalluser", s.middleware.JwtAuthentication(userHanlder.GetAllUsers))
	// route.GET("/listalluser", userHanlder.GetAllUsers)

	go func() {
		for {
			time.Sleep(time.Second * 10)
			count, _ := userUseCase.CountUser(context.Background())
			log.Printf("users in database is %d", count)
		}
	}()
	// route.GET("/countuser", userHanlder.CountUser)
	route.GET("/getuser", s.middleware.JwtAuthentication(userHanlder.GetUserById))

	route.PATCH("/updateuser", s.middleware.JwtAuthentication(userHanlder.UpdateUserDetail))
	route.DELETE("/deleteuser", s.middleware.JwtAuthentication(userHanlder.DeleteUser))

}
