package repositories

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/guatom999/backend-challenge/modules"
	"github.com/guatom999/backend-challenge/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	RepositoryInterface interface {
		IsUserAlreadyExist(pctx context.Context, email string) bool
		CreateUser(pctx context.Context, user *modules.User) error
		GetAllUser(pctx context.Context) ([]*modules.ListUserRes, error)
		GetUserById(pctx context.Context, Id string) (*modules.ListUserRes, error)
		UpdateUser(pctx context.Context, Id string, userUpdateReq *modules.UpdateUser) error
		DeleteUser(pctx context.Context, Id string) error
	}

	repository struct {
		db *mongo.Client
	}
)

func NewRepository(db *mongo.Client) RepositoryInterface {
	return &repository{db: db}
}

func (r *repository) IsUserAlreadyExist(pctx context.Context, email string) bool {

	ctx, cancel := context.WithTimeout(pctx, time.Second*5)
	defer cancel()

	db := r.db.Database("user_db")
	collection := db.Collection("users")

	result := new(modules.User)

	if err := collection.FindOne(ctx, bson.M{"email": email}).Decode(result); err != nil {
		return false
	}

	log.Printf("Error: User Already Exist")

	return true

}

func (r *repository) CreateUser(pctx context.Context, user *modules.User) error {

	ctx, cancel := context.WithTimeout(pctx, time.Second*5)
	defer cancel()

	db := r.db.Database("user_db")
	collection := db.Collection("users")

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Printf("Error: Create User Failed %s", err.Error())
		return errors.New("error: failed to create user")
	}

	return nil

}

func (r *repository) GetAllUser(pctx context.Context) ([]*modules.ListUserRes, error) {

	ctx, cancel := context.WithTimeout(pctx, time.Second*5)
	defer cancel()

	db := r.db.Database("user_db")
	collection := db.Collection("users")

	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Error: Get All User Failed %s", err.Error())
		return make([]*modules.ListUserRes, 0), errors.New("error: failed to get all user")
	}

	results := make([]*modules.ListUserRes, 0)

	for cur.Next(ctx) {
		result := new(modules.User)

		if err := cur.Decode(result); err != nil {
			log.Printf("Error: Get All User Failed %s", err.Error())
			return make([]*modules.ListUserRes, 0), errors.New("error: failed to get all user")
		}

		results = append(results, &modules.ListUserRes{
			ID:    result.ID.Hex(),
			Name:  result.Name,
			Email: result.Email,
		})
	}

	return results, nil

}

func (r *repository) GetUserById(pctx context.Context, Id string) (*modules.ListUserRes, error) {

	ctx, cancel := context.WithTimeout(pctx, time.Second*10)
	defer cancel()

	db := r.db.Database("user_db")
	collection := db.Collection("users")

	result := new(modules.User)

	if err := collection.FindOne(ctx, bson.M{"_id": utils.ConvertStringToPrimitiveId(Id)}).Decode(result); err != nil {
		log.Printf("Error: Get User By Id Failed %s", err.Error())
		return nil, errors.New("errros: failed to get user by id")
	}

	return &modules.ListUserRes{
		ID:    result.ID.Hex(),
		Name:  result.Name,
		Email: result.Email,
	}, nil

}

func (r *repository) UpdateUser(pctx context.Context, Id string, userUpdateReq *modules.UpdateUser) error {

	ctx, cancel := context.WithTimeout(pctx, time.Second*10)
	defer cancel()

	db := r.db.Database("user_db")
	collection := db.Collection("users")

	_, err := collection.UpdateOne(ctx, bson.M{"_id": utils.ConvertStringToPrimitiveId(Id)}, bson.M{"$set": userUpdateReq})
	if err != nil {
		log.Printf("Error: Update User Failed %s", err.Error())
		return errors.New("error: failed to update user")
	}

	return nil

}

func (r *repository) DeleteUser(pctx context.Context, Id string) error {

	ctx, cancel := context.WithTimeout(pctx, time.Second*10)
	defer cancel()

	db := r.db.Database("user_db")
	collection := db.Collection("users")

	if _, err := collection.UpdateOne(ctx, bson.M{"_id": utils.ConvertStringToPrimitiveId(Id)}, bson.M{"$set": bson.M{"is_deleted": true}}); err != nil {
		log.Printf("Error: Delete User Failed %s", err.Error())
		return errors.New("error: failed to delete user")
	}

	return nil

}
