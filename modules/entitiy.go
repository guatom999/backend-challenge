package modules

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	User struct {
		ID        primitive.ObjectID `bson:"_id,omitempty"`
		Name      string             `bson:"name"`
		Email     string             `bson:"email,unique"`
		Password  string             `bson:"password"`
		IsDeleted bool               `bson:"is_deleted"`
		CreatedAt time.Time          `bson:"created_at"`
	}

	UpdateUser struct {
		Name  string `bson:"name"`
		Email string `bson:"email"`
	}
)
