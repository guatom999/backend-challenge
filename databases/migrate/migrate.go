package main

import (
	"context"
	"log"

	"github.com/guatom999/backend-challenge/config"
	"github.com/guatom999/backend-challenge/databases"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ()

func main() {
	ctx := context.Background()

	cfg := config.GetMigrateConfig()

	db := databases.DbConnect(ctx, cfg).Database("user_db")
	defer db.Client().Disconnect(ctx)

	col := db.Collection("users")

	indexs, err := col.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys:    bson.M{"email": 1},
			Options: options.Index().SetUnique(true),
		},
	})

	if err != nil {
		panic(err)
	}

	for i, v := range indexs {
		log.Printf("Index %d is %s", i, v)
	}

	log.Println("Migrate Success")

}
