package server

import (
	"github.com/guatom999/backend-challenge/modules/handlers"
	"github.com/guatom999/backend-challenge/modules/repositories"
	"github.com/guatom999/backend-challenge/modules/usecases"
)

func (s *server) UserService() {

	userRepository := repositories.NewRepository(s.db)
	userUseCase := usecases.NewUseCase(userRepository)
	userHanlder := handlers.NewHandler(userUseCase)

	route := s.app.Group("/user")

	route.POST("/register", userHanlder.Register)

	route.GET("/listalluser", userHanlder.GetAllUsers)
	route.GET("/getuser", userHanlder.GetUserById)

	route.PATCH("/updateuser", userHanlder.UpdateUserDetail)
	route.DELETE("/deleteuser", userHanlder.DeleteUser)

}
