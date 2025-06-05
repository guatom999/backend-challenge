package test

import (
	"context"
	"errors"
	"testing"

	"github.com/guatom999/backend-challenge/modules/users"
	"github.com/guatom999/backend-challenge/modules/users/repositories"
	"github.com/guatom999/backend-challenge/modules/users/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type (
// 	testRegister struct {
// 		ctx      context.Context
// 		cfg      *config.Config
// 		req      *users.CreateUserReq
// 		expected *users.CreateUserRes
// 		isError  bool
// 	}
// )

func TestRegisterSuccess(t *testing.T) {

	cfg := NewTestConfig()
	ctx := context.Background()

	repoMock := new(repositories.RepositoryMock)
	usecase := usecases.NewUseCase(cfg, repoMock)

	repoMock.On("IsUserAlreadyExist", ctx, mock.AnythingOfType("string")).Return(false)

	repoMock.On("CreateUser", ctx, mock.AnythingOfType("*users.User")).Return(primitive.NewObjectID(), nil)

	type (
		args struct {
			label    string
			in       *users.CreateUserReq
			expected *users.CreateUserRes
		}
	)

	cases := []args{
		{
			label: "Register Success",
			in: &users.CreateUserReq{
				Name:     "TestToo",
				Email:    "TestToo@hotmail.com",
				Password: "TestToo",
			},
			expected: &users.CreateUserRes{
				ID: primitive.NewObjectID().Hex(),
			},
		},
	}

	for _, c := range cases {
		t.Run(c.label, func(t *testing.T) {
			result, err := usecase.Register(ctx, c.in)
			assert.NoError(t, err)
			assert.NotNil(t, result)
		})
	}

}

func TestRegisterFailUserAlreadyExist(t *testing.T) {

	ctx := context.Background()
	cfg := NewTestConfig()
	repoMock := new(repositories.RepositoryMock)
	usecase := usecases.NewUseCase(cfg, repoMock)

	repoMock.On("IsUserAlreadyExist", ctx, mock.AnythingOfType("string")).Return(true)

	type (
		args struct {
			label    string
			in       *users.CreateUserReq
			expected *users.CreateUserRes
		}
	)

	cases := []args{
		{
			label: "Test User Already Exist",
			in: &users.CreateUserReq{
				Name:     "TestAlreadyExist",
				Email:    "TestAlreadyExist@hotmail.com",
				Password: "TestAlreadyExist",
			},
			expected: &users.CreateUserRes{},
		},
	}

	t.Run(cases[0].label, func(t *testing.T) {
		result, err := usecase.Register(ctx, cases[0].in)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, "error: user already exist")
	})

}

func TestRegisterFail(t *testing.T) {

	ctx := context.Background()
	cfg := NewTestConfig()
	repoMock := new(repositories.RepositoryMock)
	usecase := usecases.NewUseCase(cfg, repoMock)

	repoMock.On("IsUserAlreadyExist", ctx, mock.AnythingOfType("string")).Return(false)

	repoMock.On("CreateUser", ctx, mock.AnythingOfType("*users.User")).Return(primitive.NilObjectID, errors.New("error: failed to create user"))

	type (
		args struct {
			label    string
			in       *users.CreateUserReq
			expected *users.CreateUserRes
		}
	)

	cases := args{
		label: "Case Creat User Error",
		in: &users.CreateUserReq{
			Name:     "TestAlreadyExist",
			Email:    "TestAlreadyExist@hotmail.com",
			Password: "TestAlreadyExist",
		},
		expected: &users.CreateUserRes{
			ID: primitive.NilObjectID.Hex(),
		},
	}

	t.Run(cases.label, func(t *testing.T) {
		result, err := usecase.Register(ctx, cases.in)
		assert.Error(t, err)
		assert.EqualError(t, err, "error: failed to create user")
		assert.EqualValues(t, cases.expected.ID, result.ID)
	})

}

// func TestRegisterSuccess(t *testing.T) {

// 	cfg := NewTestConfig()
// 	ctx := context.Background()

// 	repoMock := new(repositories.RepositoryMock)
// 	usecase := usecases.NewUseCase(cfg, repoMock)

// 	repoMock.On("IsUserAlreadyExist", ctx, mock.AnythingOfType("string")).Return(
// 		false,
// 	)

// 	repoMock.On("CreateUser", ctx, &users.User{
// 		Name:      mock.AnythingOfType("string"),
// 		Email:     "TestZeroOne@hotmail.com",
// 		Password:  string("TestZeroOne"),
// 		IsDeleted: false,
// 		CreatedAt: utils.GetLocalBkkTime(),
// 	}).Return(nil)

// 	tests := []testRegister{
// 		{
// 			ctx: ctx,
// 			cfg: cfg,
// 			req: &users.CreateUserReq{
// 				Name:     "TestZeroOne",
// 				Email:    "TestZeroOne@hotmail.com",
// 				Password: string("TestZeroOne"),
// 			},
// 			expected: &users.CreateUserRes{
// 				ID: primitive.NewObjectID().Hex(),
// 			},
// 			isError: false,
// 		},
// 		{
// 			ctx: ctx,
// 			cfg: cfg,
// 			req: &users.CreateUserReq{
// 				Name:     "",
// 				Email:    "TestZeroOne@hotmail.com",
// 				Password: string("TestZeroOne"),
// 			},
// 			expected: &users.CreateUserRes{
// 				ID: primitive.NilObjectID.Hex(),
// 			},
// 			isError: true,
// 		},
// 		{
// 			ctx: ctx,
// 			cfg: cfg,
// 			req: &users.CreateUserReq{
// 				Name:     "TestZeroOne",
// 				Email:    "",
// 				Password: string("TestZeroOne"),
// 			},
// 			expected: &users.CreateUserRes{
// 				ID: primitive.NilObjectID.Hex(),
// 			},
// 			isError: true,
// 		},
// 		{
// 			ctx: ctx,
// 			cfg: cfg,
// 			req: &users.CreateUserReq{
// 				Name:     "TestZeroOne",
// 				Email:    "TestZeroOne@hotmail.com",
// 				Password: string(""),
// 			},
// 			expected: &users.CreateUserRes{
// 				ID: primitive.NilObjectID.Hex(),
// 			},
// 			isError: true,
// 		},
// 	}

// 	for i, test := range tests {
// 		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {

// 			fmt.Printf("case %d", i)

// 			res, err := usecase.Register(test.ctx, test.req)
// 			if test.isError {
// 				assert.Nil(t, err)
// 			} else {
// 				assert.NoError(t, err)
// 				assert.NotNil(t, res)
// 			}
// 		})
// 	}

// }
