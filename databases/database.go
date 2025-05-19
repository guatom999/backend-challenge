package databases

import (
	"context"
	"log"
	"time"

	"github.com/guatom999/backend-challenge/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func DbConnect(pctx context.Context, cfg *config.Config) *mongo.Client {

	ctx, cancel := context.WithTimeout(pctx, time.Second*10)
	defer cancel()

	cli, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Db.Uri))
	if err != nil {
		log.Fatal("Error Failed to Connect DataBase")
		panic(err)
	}

	if err := cli.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("Error Failed to Ping DataBase")
		panic(err)
	}

	return cli

}
