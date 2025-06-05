package repositories

import (
	"context"

	"github.com/guatom999/backend-challenge/modules/users"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RepositoryMock struct {
	mock.Mock
}

func NewUserRepositoryMock() RepositoryInterface {
	return &RepositoryMock{}
}

func (m *RepositoryMock) IsUserAlreadyExist(pctx context.Context, email string) bool {
	args := m.Called(pctx, email)
	return args.Bool(0)
}
func (m *RepositoryMock) CreateUser(pctx context.Context, user *users.User) (primitive.ObjectID, error) {
	args := m.Called(pctx, user)
	return args.Get(0).(primitive.ObjectID), args.Error(1)
}
func (m *RepositoryMock) FindUserCredential(pctx context.Context, email string) (*users.User, error) {
	args := m.Called(pctx, email)
	return args.Get(0).(*users.User), args.Error(1)
}
func (m *RepositoryMock) GetAllUser(pctx context.Context) ([]*users.ListUserRes, error) {
	args := m.Called(pctx)
	return args.Get(0).([]*users.ListUserRes), args.Error(1)
}
func (m *RepositoryMock) GetUserById(pctx context.Context, Id string) (*users.ListUserRes, error) {
	args := m.Called(pctx, Id)
	return args.Get(0).(*users.ListUserRes), args.Error(1)
}
func (m *RepositoryMock) CountUser(pctx context.Context) (int64, error) {
	args := m.Called(pctx)
	return args.Get(0).(int64), args.Error(1)
}
func (m *RepositoryMock) UpdateUser(pctx context.Context, Id string, userUpdateReq *users.UpdateUser) error {
	args := m.Called(pctx, Id, userUpdateReq)
	return args.Error(0)
}
func (m *RepositoryMock) DeleteUser(pctx context.Context, Id string) error {
	args := m.Called(pctx, Id)
	return args.Error(0)
}
