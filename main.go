package main

import (
	"context"

	"github.com/guatom999/backend-challenge/config"
	"github.com/guatom999/backend-challenge/databases"
	"github.com/guatom999/backend-challenge/server"
)

func main() {

	ctx := context.Background()

	cfg := config.GetConfig()

	db := databases.DbConnect(ctx, cfg)

	server.NewEchoServer(db, cfg).Start(ctx)

}
