package usecases

import (
	"context"
	"errors"
	"log"

	"github.com/guatom999/backend-challenge/modules"
	"github.com/guatom999/backend-challenge/modules/repositories"
	"github.com/guatom999/backend-challenge/utils"
	"golang.org/x/crypto/bcrypt"
)

type (
	UsecaseInterface interface {
		Register(pctx context.Context, user *modules.CreateUserReq) error
		GetAllUses(pctx context.Context) ([]*modules.ListUserRes, error)
		GetUserById(pctx context.Context, Id string) (*modules.ListUserRes, error)
		UpdateUserDetail(pctx context.Context, updateReq *modules.UpdateUserReq) error
		Deleteuser(pctx context.Context, Id string) error
	}

	usecase struct {
		repository repositories.RepositoryInterface
	}
)

func NewUseCase(repository repositories.RepositoryInterface) UsecaseInterface {
	return &usecase{repository: repository}
}

func (u *usecase) Register(pctx context.Context, user *modules.CreateUserReq) error {

	if u.repository.IsUserAlreadyExist(pctx, user.Email) {
		return errors.New("error: user already exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Printf("Error: Failed to hash password %s", err.Error())
		return err
	}

	if err := u.repository.CreateUser(pctx, &modules.User{
		Name:      user.Name,
		Email:     user.Email,
		Password:  string(hashedPassword),
		IsDeleted: false,
		CreatedAt: utils.GetLocalBkkTime(),
	}); err != nil {
		return err
	}

	return nil

}

func (u *usecase) GetAllUses(pctx context.Context) ([]*modules.ListUserRes, error) {

	users, err := u.repository.GetAllUser(pctx)
	if err != nil {
		return users, err
	}

	return users, nil

}

func (u *usecase) GetUserById(pctx context.Context, Id string) (*modules.ListUserRes, error) {

	user, err := u.repository.GetUserById(pctx, Id)
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (u *usecase) UpdateUserDetail(pctx context.Context, updateReq *modules.UpdateUserReq) error {

	if err := u.repository.UpdateUser(pctx, updateReq.ID, &modules.UpdateUser{
		Name:  updateReq.Name,
		Email: updateReq.Email,
	}); err != nil {
		return err
	}

	return nil
}

func (u *usecase) Deleteuser(pctx context.Context, Id string) error {

	if err := u.repository.DeleteUser(pctx, Id); err != nil {
		return err
	}

	return nil

}
