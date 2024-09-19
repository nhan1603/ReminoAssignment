package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/nhan1603/ReminoAssignment/api/internal/appconfig/db/pg"
	"github.com/nhan1603/ReminoAssignment/api/internal/appconfig/httpserver"
	"github.com/nhan1603/ReminoAssignment/api/internal/controller/auth"
	"github.com/nhan1603/ReminoAssignment/api/internal/repository"
	"github.com/nhan1603/ReminoAssignment/api/internal/repository/generator"
)

func main() {
	log.Println("Remitano Assignment API")

	ctx := context.Background()

	// Initial DB connection
	conn, err := pg.Connect(os.Getenv("PG_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Initial snowflake generator
	if err := generator.InitSnowflakeGenerators(); err != nil {
		log.Fatal(err)
	}

	// Initial router
	rtr, err := initRouter(ctx, conn)
	if err != nil {
		log.Fatal(err)
	}

	// Start server
	httpserver.Start(httpserver.Handler(ctx, rtr.routes))
}

func initRouter(
	ctx context.Context,
	db *sql.DB,
) (router, error) {
	repo := repository.New(db)

	return router{
		ctx:      ctx,
		authCtrl: auth.New(repo),
	}, nil
}
