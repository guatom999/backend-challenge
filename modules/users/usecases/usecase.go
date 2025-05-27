package usecases

import (
	"context"
	"errors"
	"log"

	"github.com/guatom999/backend-challenge/config"
	"github.com/guatom999/backend-challenge/modules/users"
	"github.com/guatom999/backend-challenge/modules/users/repositories"
	"github.com/guatom999/backend-challenge/pkg/jwtauth"
	"github.com/guatom999/backend-challenge/utils"
	"golang.org/x/crypto/bcrypt"
)

type (
	UsecaseInterface interface {
		Register(pctx context.Context, user *users.CreateUserReq) error
		GetAllUses(pctx context.Context) ([]*users.ListUserRes, error)
		CountUser(pctx context.Context) (int64, error)
		GetUserById(pctx context.Context, Id string) (*users.ListUserRes, error)
		Login(pctx context.Context, user *users.LoginCredentialReq) (*users.LoginCredentialRes, error)
		UpdateUserDetail(pctx context.Context, userId string, updateReq *users.UpdateUserReq) error
		Deleteuser(pctx context.Context, Id string) error
	}

	usecase struct {
		cfg        *config.Config
		repository repositories.RepositoryInterface
	}
)

func NewUseCase(cfg *config.Config, repository repositories.RepositoryInterface) UsecaseInterface {
	return &usecase{cfg: cfg, repository: repository}
}

func (u *usecase) Register(pctx context.Context, user *users.CreateUserReq) error {

	if u.repository.IsUserAlreadyExist(pctx, user.Email) {
		return errors.New("error: user already exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Printf("Error: Failed to hash password %s", err.Error())
		return err
	}

	if err := u.repository.CreateUser(pctx, &users.User{
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

func (u *usecase) Login(pctx context.Context, user *users.LoginCredentialReq) (*users.LoginCredentialRes, error) {

	userCredential, err := u.repository.FindUserCredential(pctx, user.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userCredential.Password), []byte(user.Password)); err != nil {
		log.Printf("Error: Failed to compare password %s", err.Error())
		return nil, errors.New("error:invalid email or password")
	}

	claims := jwtauth.NewJwtToken(u.cfg.Jwt.Secret, &jwtauth.Claims{
		UserId: userCredential.ID.Hex(),
	}).SignToken()

	// claims := users.AuthClaims{
	// 	Claims: &users.Claims{
	// 		UserId: userCredential.ID.Hex(),
	// 	},
	// 	RegisteredClaims: jwt.RegisteredClaims{
	// 		Issuer:    "user-challenge",
	// 		Subject:   "jwt-token",
	// 		Audience:  []string{"user"},
	// 		ExpiresAt: jwt.NewNumericDate(utils.GetLocalBkkTime().Add(time.Second * 60)),
	// 		NotBefore: jwt.NewNumericDate(utils.GetLocalBkkTime()),
	// 		IssuedAt:  jwt.NewNumericDate(utils.GetLocalBkkTime()),
	// 	},
	// }

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// jwtToken, err := token.SignedString([]byte(u.cfg.Jwt.Secret))
	// if err != nil {
	// 	log.Printf("Error: Failed to sign token %s", err.Error())
	// 	return nil, errors.New("error:something went wrong")
	// }

	return &users.LoginCredentialRes{
		UserId: userCredential.ID.Hex(),
		Token:  claims,
	}, nil
}

func (u *usecase) GetAllUses(pctx context.Context) ([]*users.ListUserRes, error) {

	users, err := u.repository.GetAllUser(pctx)
	if err != nil {
		return users, err
	}

	return users, nil

}

func (u *usecase) CountUser(pctx context.Context) (int64, error) {

	return u.repository.CountUser(pctx)

}

func (u *usecase) GetUserById(pctx context.Context, Id string) (*users.ListUserRes, error) {

	user, err := u.repository.GetUserById(pctx, Id)
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (u *usecase) UpdateUserDetail(pctx context.Context, userId string, updateReq *users.UpdateUserReq) error {

	if err := u.repository.UpdateUser(pctx, userId, &users.UpdateUser{
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
