package middlewareRepositories

import "go.mongodb.org/mongo-driver/mongo"

type (
	MiddlewareRepositoryInterface interface {
	}

	middlewareRepository struct {
		db *mongo.Client
	}
)

func NewRepository(db *mongo.Client) MiddlewareRepositoryInterface {
	return &middlewareRepository{
		db: db,
	}
}
